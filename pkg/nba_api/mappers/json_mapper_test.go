package mappers

import (
	"encoding/json"
	"os"
	"testing"
)

func readJSONFileAsRawMessage(path string) (json.RawMessage, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b), nil
}

func TestBuildResultSets(t *testing.T) {
	raw, err := readJSONFileAsRawMessage("mock/mock.json")
	if err != nil {
		t.Fatal(err)
	}

	res, err := MapResultSetsToResponseMap(raw)
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatal("res is nil")
	}

	tsAny, ok := res["TeamStats"]
	if !ok {
		t.Fatalf(`missing key "TeamStats"`)
	}
	tsRows, ok := tsAny.([]map[string]any)
	if !ok {
		t.Fatalf(`"TeamStats" has unexpected type %T`, tsAny)
	}
	if len(tsRows) != 2 {
		t.Fatalf(`"TeamStats" expected 2 rows, got %d`, len(tsRows))
	}

	if got := tsRows[0]["TEAM_ID"]; got != float64(1610612747) {
		t.Fatalf(`TeamStats[0]["TEAM_ID"] = %#v (%T), want %v`, got, got, float64(1610612747))
	}
	if got := tsRows[0]["TEAM_NAME"]; got != "Lakers" {
		t.Fatalf(`TeamStats[0]["TEAM_NAME"] = %#v (%T), want %q`, got, got, "Lakers")
	}
	if got := tsRows[0]["PTS"]; got != float64(112) {
		t.Fatalf(`TeamStats[0]["PTS"] = %#v (%T), want %v`, got, got, float64(112))
	}
	if got := tsRows[0]["AST"]; got != float64(26) {
		t.Fatalf(`TeamStats[0]["AST"] = %#v (%T), want %v`, got, got, float64(26))
	}
	if got := tsRows[0]["WINNER"]; got != true {
		t.Fatalf(`TeamStats[0]["WINNER"] = %#v (%T), want %v`, got, got, true)
	}

	if got := tsRows[1]["WINNER"]; got != nil {
		t.Fatalf(`TeamStats[1]["WINNER"] = %#v (%T), want nil`, got, got)
	}

	pAny, ok := res["Players"]
	if !ok {
		t.Fatalf(`missing key "Players"`)
	}
	pRows, ok := pAny.([]map[string]any)
	if !ok {
		t.Fatalf(`"Players" has unexpected type %T`, pAny)
	}
	if len(pRows) != 2 {
		t.Fatalf(`"Players" expected 2 rows, got %d`, len(pRows))
	}
	if got := pRows[0]["PLAYER_NAME"]; got != "LeBron James" {
		t.Fatalf(`Players[0]["PLAYER_NAME"] = %#v (%T), want %q`, got, got, "LeBron James")
	}

	giAny, ok := res["GameInfo"]
	if !ok {
		t.Fatalf(`missing key "GameInfo"`)
	}
	giRows, ok := giAny.([]map[string]any)
	if !ok {
		t.Fatalf(`"GameInfo" has unexpected type %T`, giAny)
	}
	if len(giRows) != 1 {
		t.Fatalf(`"GameInfo" expected 1 row, got %d`, len(giRows))
	}
	if got := giRows[0]["GAME_ID"]; got != "0022300001" {
		t.Fatalf(`GameInfo[0]["GAME_ID"] = %#v (%T), want %q`, got, got, "0022300001")
	}
	if got := giRows[0]["ATTENDANCE"]; got != float64(18997) {
		t.Fatalf(`GameInfo[0]["ATTENDANCE"] = %#v (%T), want %v`, got, got, float64(18997))
	}
}
