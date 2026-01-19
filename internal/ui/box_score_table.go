package ui

import (
	"fmt"
	"nbarecap/pkg/nba_api/models"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
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

func appendTeamIdRow(rows []table.Row, boxScore *models.BoxScoreTraditionalV3) []table.Row {
	rows = append(rows, table.Row{"", "", "", "", ""})
	rows = append(rows, table.Row{
		fmt.Sprintf("— %s (%s) —", boxScore.HomeTeam.TeamName, boxScore.HomeTeam.TeamTricode),
		"", "", "",
	})
	return rows
}

func boxScoreToRows(boxScore *models.BoxScoreTraditionalV3) []table.Row {
	homePlayers := boxScore.HomeTeam.Players
	awayPlayers := boxScore.AwayTeam.Players

	rows := make([]table.Row, 0, len(homePlayers)+len(awayPlayers))

	rows = appendTeamIdRow(rows, boxScore)
	for _, p := range boxScore.HomeTeam.Players {
		rows = append(rows, playerStatsToRow(&p))
	}

	rows = appendTeamIdRow(rows, boxScore)
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
	/* TODO: table styles & colors */
	t.SetStyles(table.DefaultStyles())
	return t
}

func (m model) renderBoxScoreView(header string) string {
	if m.currentBoxScore == nil {
		return lipgloss.Place(
			m.termWidth,
			m.termHeight,
			lipgloss.Center,
			lipgloss.Center,
			header+boxScoreLoadingMsg,
		)
	}

	boxView := tableBoxStyle.Render(m.boxTable.View())

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		header+"\n"+boxView+boxScoreLoadedMsg,
	)
}
