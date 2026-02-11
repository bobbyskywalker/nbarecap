package nba

import "testing"

func TestGetStandingsForSeason(t *testing.T) {
	items, err := GetStandingsForSeason("2024-25")
	if err != nil {
		t.Errorf("Error fetching standings: %v", err)
	}
	if items == nil {
		t.Errorf("Expected non-nil standings object, got nil")
	}
}
