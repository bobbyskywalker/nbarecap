package nba

import "testing"

func TestGetBoxScoreForGame(t *testing.T) {
	res, err := GetBoxScoreForGame("0022000001")
	if err != nil {
		t.Errorf("Error fetching box score: %v", err)
	}
	if res == nil {
		t.Errorf("Expected non-nil box score, got nil")
	}
}
