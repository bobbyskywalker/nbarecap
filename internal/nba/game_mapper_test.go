package nba

import (
	"nbarecap/internal/models"
	"testing"
)

var mockGameHeaderRows = []map[string]any{
	{
		HGameId:     "001",
		HStatusText: "Final",
		HGamecode:   "20250101" + SplitSep + "LALBOS",
		HArenaName:  "Boston Garden",
		HNatlTv:     "ESPN",
		HHomeTv:     "NBCSB",
		HAwayTv:     "SpecSN",
		HHomeTeamId: "BOS",
		HVisTeamId:  "LAL",
	},
	{
		HGameId:     "002",
		HStatusText: "Scheduled",
		HGamecode:   "20250102" + SplitSep + "NYKMIA",
		HArenaName:  "TheCoolArena",
		HNatlTv:     "",
		HHomeTv:     "BSSUN",
		HAwayTv:     "MSG",
		HHomeTeamId: "MIA",
		HVisTeamId:  "NYK",
	},
}

var mockResponseMap = map[string]any{
	RsLinescore: []map[string]any{
		{
			HGameId:         "001",
			HTeamId:         "BOS",
			HTeamAbbr:       "BOS",
			HTeamWinsLosses: "30-10",
			HPts:            117,
		},
		{
			HGameId:         "001",
			HTeamId:         "LAL",
			HTeamAbbr:       "LAL",
			HTeamWinsLosses: "22-18",
			HPts:            110,
		},
		{
			HGameId:         "002",
			HTeamId:         "MIA",
			HTeamAbbr:       "MIA",
			HTeamWinsLosses: "18-20",
			HPts:            nil,
		},
		{
			HGameId:         "002",
			HTeamId:         "NYK",
			HTeamAbbr:       "NYK",
			HTeamWinsLosses: "25-15",
			HPts:            nil,
		},
	},
}

func checkGhRowsState(rows map[string]*models.GameInfo) bool {
	if len(rows) != 2 {
		return true
	}

	g1, ok := rows["001"]
	if !ok || g1 == nil {
		return true
	}

	if g1.GameID != "001" ||
		g1.Status != "Final" ||
		g1.Arena != "Boston Garden" ||
		g1.NatTV != "ESPN" ||
		g1.HomeTV != "NBCSB" ||
		g1.AwayTV != "SpecSN" ||
		g1.HomeID != "BOS" ||
		g1.AwayID != "LAL" {
		return true
	}

	if g1.AwayAbbr != "LAL" || g1.HomeAbbr != "BOS" {
		return true
	}

	expectedSortKey1 := "Final" + SortKeySep + "001"
	if g1.SortKey != expectedSortKey1 {
		return true
	}

	g2, ok := rows["002"]
	if !ok || g2 == nil {
		return true
	}

	if g2.GameID != "002" ||
		g2.Status != "Scheduled" ||
		g2.Arena != "TheCoolArena" ||
		g2.HomeTV != "BSSUN" ||
		g2.AwayTV != "MSG" ||
		g2.HomeID != "MIA" ||
		g2.AwayID != "NYK" {
		return true
	}

	if g2.AwayAbbr != "NYK" || g2.HomeAbbr != "MIA" {
		return true
	}

	expectedSortKey2 := "Scheduled" + SortKeySep + "002"
	if g2.SortKey != expectedSortKey2 {
		return true
	}

	return false
}

func checkLineScoreRowsState(rows map[string]*models.GameInfo) bool {
	g1 := rows["001"]
	if g1 == nil {
		return true
	}

	if g1.Home.Abbr != "BOS" ||
		g1.Home.Record != "30-10" ||
		g1.Home.Pts == nil ||
		*g1.Home.Pts != 117 {
		return true
	}

	if g1.Away.Abbr != "LAL" ||
		g1.Away.Record != "22-18" ||
		g1.Away.Pts == nil ||
		*g1.Away.Pts != 110 {
		return true
	}

	g2 := rows["002"]
	if g2 == nil {
		return true
	}

	if g2.Home.Abbr != "MIA" ||
		g2.Home.Record != "18-20" ||
		g2.Home.Pts != nil {
		return true
	}

	if g2.Away.Abbr != "NYK" ||
		g2.Away.Record != "25-15" ||
		g2.Away.Pts != nil {
		return true
	}

	return false
}

func TestBuildGameMap(t *testing.T) {
	gameMap := buildGameMap(mockResponseMap, mockGameHeaderRows)
	if checkGhRowsState(gameMap) || checkLineScoreRowsState(gameMap) {
		t.Errorf("Game map state is not as expected")
	}
}
