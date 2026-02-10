package mappers

import (
	"encoding/json"
	"fmt"

	"nbarecap/pkg/nba_api/models"
)

func BuildStandingsList(responseMap map[string]any) ([]*models.Standing, error) {
	standingsList := make([]*models.Standing, 0)

	standingsRows, ok := responseMap["Standings"]
	if !ok {
		return nil, fmt.Errorf("error mapping standings")
	}

	rows, ok := standingsRows.([]map[string]any)
	if !ok {
		return nil, fmt.Errorf("error mapping standings")
	}

	for _, row := range rows {
		bytes, err := json.Marshal(row)
		if err != nil {
			continue
		}

		var s models.Standing
		if err := json.Unmarshal(bytes, &s); err != nil {
			continue
		}
		standingsList = append(standingsList, &s)
	}

	return standingsList, nil
}
