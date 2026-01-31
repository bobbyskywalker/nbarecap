package ui

import (
	"fmt"
	"nbarecap/internal/nba"
	"nbarecap/pkg/nba_api/models"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	preGameStatus      = "PRE-GAME"
	boxScoreManual     = "<-/-> switch teams"
	statModeManual     = "'m' switch stats mode"
	goBackManual       = "'esc' go back"
	percentageFormat   = "%.1f%s"
	boxScoreLoadingMsg = "\nLoading box score...\n\n(esc to go back)"
)

type boxScoreMsg struct {
	score *models.BoxScoreTraditionalV3
	err   error
}

func (m appModel) fetchBoxScoreCmd(gameID string) tea.Cmd {
	return func() tea.Msg {
		bx, err := nba.GetBoxScoreForGame(gameID)
		if err != nil {
			return boxScoreMsg{err: err}
		}
		return boxScoreMsg{score: bx}
	}
}

func playerStatsToRow(player *models.PlayerV3, showingPercentages bool) table.Row {
	playerStats := player.Statistics
	commentOrMinutes := playerStats.Minutes
	if playerStats.Minutes == "" {
		commentOrMinutes = strings.Split(player.Comment, " ")[0]
	}

	if !showingPercentages {
		return table.Row{
			player.NameI,
			commentOrMinutes,
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
	return table.Row{
		player.NameI,
		strconv.Itoa(playerStats.FieldGoalsMade),
		strconv.Itoa(playerStats.FieldGoalsAttempted),
		fmt.Sprintf(percentageFormat, playerStats.FieldGoalsPercentage*100, "%"),
		strconv.Itoa(playerStats.ThreePointersMade),
		strconv.Itoa(playerStats.ThreePointersAttempted),
		fmt.Sprintf(percentageFormat, playerStats.ThreePointersPercentage*100, "%"),
		strconv.Itoa(playerStats.FreeThrowsMade),
		strconv.Itoa(playerStats.FreeThrowsAttempted),
		fmt.Sprintf(percentageFormat, playerStats.FreeThrowsPercentage*100, "%"),
	}
}

func appendTeamIdRow(rows []table.Row, team models.TeamV3) []table.Row {
	color := TeamColors[team.TeamTricode]

	rows = append(rows, table.Row{"", "", "", "", ""})
	rows = append(rows, table.Row{
		lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(
			fmt.Sprintf("— %s (%s) —", team.TeamName, team.TeamTricode),
			"", "", ""),
	})
	return rows
}

func boxScoreToRows(boxScore *models.BoxScoreTraditionalV3, showingAway bool, showingPercentages bool) []table.Row {
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
		rows = append(rows, playerStatsToRow(&p, showingPercentages))
	}
	return rows
}

func resolveColumnNames(showingPercentages bool) []table.Column {
	if !showingPercentages {
		return []table.Column{
			{Title: "Player", Width: 40},
			{Title: "MIN", Width: 10},
			{Title: "PTS", Width: 5},
			{Title: "REB", Width: 5},
			{Title: "AST", Width: 5},
			{Title: "STL", Width: 5},
			{Title: "BLK", Width: 5},
			{Title: "TOV", Width: 5},
			{Title: "PF", Width: 5},
			{Title: "+/-", Width: 5},
		}
	}
	return []table.Column{
		{Title: "Player", Width: 40},
		{Title: "FGM", Width: 5},
		{Title: "FGA", Width: 5},
		{Title: "FG%", Width: 7},
		{Title: "3PM", Width: 5},
		{Title: "3PA", Width: 5},
		{Title: "3P%", Width: 7},
		{Title: "FTM", Width: 5},
		{Title: "FTA", Width: 5},
		{Title: "FT%", Width: 7},
	}
}

func newBoxScoreTable(boxScore *models.BoxScoreTraditionalV3, showingAway bool, showingPercentages bool) table.Model {
	columns := resolveColumnNames(showingPercentages)
	rows := boxScoreToRows(boxScore, showingAway, showingPercentages)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

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

func buildBoxScoreHeader(m appModel) string {
	homeTricode := m.currentBoxScore.HomeTeam.TeamTricode
	awayTricode := m.currentBoxScore.AwayTeam.TeamTricode

	homePts := m.currentBoxScore.HomeTeam.Statistics.Points
	awayPts := m.currentBoxScore.AwayTeam.Statistics.Points

	away := createBigTeamBadgeStyle(awayTricode).Render(awayTricode)
	home := createBigTeamBadgeStyle(homeTricode).Render(homeTricode)

	scoreText := strconv.Itoa(awayPts) + "  —  " + strconv.Itoa(homePts)
	score := boxScoreHeaderScoreStyle.Render(scoreText)

	scoreHeader := boxScoreHeaderRowStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Center, away, "  ", score, "  ", home),
	)

	statusText := preGameStatus
	if homePts != awayPts != false {
		statusText = "FINAL"
	}

	statusHeader := boxScoreHeaderStatusPillStyle.Render(strings.ToUpper(statusText))

	dateHeader := boxScoreHeaderDateStyle.Render(m.date.Format(dateFormat))
	return lipgloss.JoinVertical(lipgloss.Center, scoreHeader, statusHeader, dateHeader)
}

func buildBoxScoreFooter(m appModel) string {
	var dots string
	if !m.showingAway {
		dots = dotInactiveStyle.Render(dotInactiveIcon) + dotActiveStyle.Render(dotActiveIcon)
	} else {
		dots = dotActiveStyle.Render(dotActiveIcon) + dotInactiveStyle.Render(dotInactiveIcon)
	}
	manuals := manualTextStyle.Render(" • " + boxScoreManual + " • " + statModeManual + " • " + goBackManual)
	return lipgloss.JoinVertical(lipgloss.Center, dots+"\n\n"+manuals)
}
