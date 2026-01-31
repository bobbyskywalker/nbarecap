package ui

import (
	"fmt"
	"nbarecap/internal/nba"
	"nbarecap/pkg/nba_api/models"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	playByPlayScoreFormat = "%s-%s"
	playByPlayLoadingMsg  = "\nLoading play by play...\n\n(esc to go back)"
	gameClockFormat       = "%s:%s"
	distanceFormat        = "%dft"
	playByPlayManual      = "<- / -> switch periods"
	quarterFormat         = "Q%d"
	overtimeFormat        = "OT%d"
)

/*
	TODO: unique styling for this view
	A refresh button for live score fetching
*/

type playByPlayMsg struct {
	content *models.PlayByPlayV3
	err     error
}

func (m appModel) fetchPlayByPlayCmd(gameID string) tea.Cmd {
	return func() tea.Msg {
		pb, err := nba.GetPlayByPlayForGame(gameID)
		if err != nil {
			return playByPlayMsg{nil, err}
		}
		return playByPlayMsg{pb, err}
	}
}

func buildPlayByPlayRows(playByPlay *models.PlayByPlayV3, period int) []table.Row {
	var rows []table.Row

	var periodStr string
	if period <= 4 {
		periodStr = fmt.Sprintf(quarterFormat, period)
	} else {
		periodStr = fmt.Sprintf(overtimeFormat, period-4)
	}

	for _, action := range playByPlay.Game.Actions {
		if action.Period == period {
			rows = append(rows, table.Row{
				periodStr,
				fmt.Sprintf(gameClockFormat, action.Clock[2:4], action.Clock[5:len(action.Clock)-1]),
				action.TeamTricode,
				action.PlayerNameI,
				action.ActionType,
				action.ShotResult,
				fmt.Sprintf(distanceFormat, action.ShotDistance),
				fmt.Sprintf(playByPlayScoreFormat, action.ScoreAway, action.ScoreHome),
			})
		}
	}
	return rows
}

func newPlayByPlayTable(playByPlay *models.PlayByPlayV3, period int) table.Model {
	columns := []table.Column{
		{Title: "Period", Width: 10},
		{Title: "Time", Width: 20},
		{Title: "Team", Width: 10},
		{Title: "Player", Width: 25},
		{Title: "Action", Width: 25},
		{Title: "Result", Width: 25},
		{Title: "Dist", Width: 10},
		{Title: "Score", Width: 10},
	}

	rows := buildPlayByPlayRows(playByPlay, period)

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

func (m appModel) renderPlayByPlayView(header string, footer string) string {
	if m.currentPlayByPlay == nil {
		return lipgloss.Place(
			m.termWidth,
			m.termHeight,
			lipgloss.Center,
			lipgloss.Center,
			playByPlayLoadingMsg,
		)
	}

	pbpView := listBoxStyle.Render(m.playByPlayTable.View())

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		header+"\n"+pbpView+"\n"+footer,
	)
}

func buildPlayByPlayScoreFooter(m appModel) string {

	var dots string
	for i := 0; i < m.maxPeriod; i++ {
		if m.selectedPeriod == i+1 {
			dots += dotActiveStyle.Render(dotActiveIcon)
		} else {
			dots += dotInactiveStyle.Render(dotInactiveIcon)
		}
	}
	manual := manualTextStyle.Render(" • " + playByPlayManual + " • " + goBackManual)
	return lipgloss.JoinVertical(lipgloss.Center, dots+"\n\n"+manual)
}

func getMaxPeriod(playByPlay *models.PlayByPlayV3) int {
	maxPeriod := 0
	for _, action := range playByPlay.Game.Actions {
		if action.Period > maxPeriod {
			maxPeriod = action.Period
		}
	}
	return maxPeriod
}
