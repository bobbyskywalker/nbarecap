package nba

import (
	"testing"
)

func TestGetGetPlayByPlayForGame(t *testing.T) {
	items, err := GetPlayByPlayForGame("0022000001")
	if err != nil {
		t.Errorf("Error fetching playByPlay: %v", err)
	}
	if items == nil {
		t.Errorf("Expected non-nil playByPlay object, got nil")
	}
}
