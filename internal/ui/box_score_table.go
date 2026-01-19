package ui

import (
	"nbarecap/pkg/nba_api/models"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

func playerStatsToRow(player *models.PlayerV3) table.Row {
	playerStats := player.Statistics
	return table.Row{
		player.NameI,
		strconv.Itoa(playerStats.Points),
		strconv.Itoa(playerStats.ReboundsTotal),
		strconv.Itoa(playerStats.Assists),
		strconv.Itoa(playerStats.Steals),
		strconv.Itoa(playerStats.Blocks),
		strconv.Itoa(playerStats.Turnovers),
		strconv.Itoa(playerStats.FoulsPersonal),
		strconv.FormatFloat(playerStats.PlusMinusPoints, 'f', 1, 64),
	}
}

func boxScoreToRows(boxScore *models.BoxScoreTraditionalV3) []table.Row {
	homePlayers := boxScore.HomeTeam.Players
	awayPlayers := boxScore.AwayTeam.Players

	rows := make([]table.Row, 0, len(homePlayers)+len(awayPlayers))

	for _, p := range boxScore.HomeTeam.Players {
		rows = append(rows, playerStatsToRow(&p))
	}
	for _, p := range boxScore.AwayTeam.Players {
		rows = append(rows, playerStatsToRow(&p))
	}
	return rows
}

func newBoxScoreTable(boxScore *models.BoxScoreTraditionalV3) table.Model {
	columns := []table.Column{
		{Title: "Player", Width: 20},
		{Title: "PTS", Width: 5},
		{Title: "REB", Width: 5},
		{Title: "AST", Width: 5},
		{Title: "STL", Width: 5},
		{Title: "BLK", Width: 5},
		{Title: "TOV", Width: 5},
		{Title: "PF", Width: 5},
		{Title: "+/-", Width: 5},
	}
	rows := boxScoreToRows(boxScore)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	/* TODO: table styles */
	t.SetStyles(table.DefaultStyles())
	return t
}
