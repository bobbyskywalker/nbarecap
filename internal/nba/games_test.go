package nba

import (
	"testing"
	"time"
)

func TestGetAllGamesForDate(t *testing.T) {
	validDateStr := "2021-01-01"
	date, _ := time.Parse("2006-01-02", validDateStr)
	items, err := GetAllGamesForDate(&date)
	if err != nil {
		t.Errorf("Error fetching games: %v", err)
	}
	if items == nil {
		t.Errorf("Expected non-nil list of games, got nil")
	}
}
