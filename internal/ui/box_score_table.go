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

func appendTeamIdRow(rows []table.Row, team models.TeamV3) []table.Row {
	rows = append(rows, table.Row{"", "", "", "", ""})
	rows = append(rows, table.Row{
		fmt.Sprintf("— %s (%s) —", team.TeamName, team.TeamTricode),
		"", "", "",
	})
	return rows
}

func boxScoreToRows(boxScore *models.BoxScoreTraditionalV3, showingAway bool) []table.Row {
	var players []models.PlayerV3
	var team models.TeamV3

	if showingAway {
		players = boxScore.AwayTeam.Players
		team = boxScore.AwayTeam
	} else {
		players = boxScore.HomeTeam.Players
		team = boxScore.HomeTeam
	}

	rows := make([]table.Row, 0, len(players)+1)
	rows = appendTeamIdRow(rows, team)
	for _, p := range players {
		rows = append(rows, playerStatsToRow(&p))
	}
	return rows
}

func newBoxScoreTable(boxScore *models.BoxScoreTraditionalV3, showingAway bool) table.Model {
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
	rows := boxScoreToRows(boxScore, showingAway)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	/* TODO: table styles & colors */
	t.SetStyles(table.Styles{
		Header:   gamesListHeaderStyle,
		Cell:     lipgloss.Style{}.Italic(true),
		Selected: selectedItemStyle,
	})
	return t
}

func (m appModel) renderBoxScoreView(header string, footer string) string {
	if m.currentBoxScore == nil {
		return lipgloss.Place(
			m.termWidth,
			m.termHeight,
			lipgloss.Center,
			lipgloss.Center,
			header+boxScoreLoadingMsg,
		)
	}

	boxView := listBoxStyle.Render(m.boxTable.View())

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		header+"\n"+boxView+"\n"+footer,
	)
}
