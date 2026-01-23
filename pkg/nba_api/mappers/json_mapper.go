package mappers

import (
	"encoding/json"
	"fmt"
)

type rsTable struct {
	Name    string   `json:"name"`
	Headers []string `json:"headers"`
	RowSet  [][]any  `json:"rowSet"`
}

type responseEnvelope struct {
	ResultSets []rsTable `json:"resultSets"`
	ResultSet  *rsTable  `json:"resultSet"`
}

func MapResultSetsToResponseMap(raw json.RawMessage) (map[string]any, error) {
	var env responseEnvelope
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, err
	}

	out := make(map[string]any, len(env.ResultSets)+1)

	addTable := func(t rsTable) error {
		rows := make([]map[string]any, 0, len(t.RowSet))
		for _, row := range t.RowSet {
			m := make(map[string]any, len(t.Headers))
			for i, h := range t.Headers {
				if i < len(row) {
					m[h] = row[i]
				} else {
					m[h] = nil
				}
			}
			rows = append(rows, m)
		}
		out[t.Name] = rows
		return nil
	}

	for _, t := range env.ResultSets {
		if t.Name == "" {
			return nil, fmt.Errorf("resultSet missing name")
		}
		if err := addTable(t); err != nil {
			return nil, err
		}
	}
	if env.ResultSet != nil && env.ResultSet.Name != "" {
		if err := addTable(*env.ResultSet); err != nil {
			return nil, err
		}
	}

	return out, nil
}
