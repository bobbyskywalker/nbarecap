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
)

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

func buildPlayByPlayRows(playByPlay *models.PlayByPlayV3) []table.Row {
	var rows []table.Row

	for _, action := range playByPlay.Game.Actions {
		rows = append(rows, table.Row{
			fmt.Sprintf(gameClockFormat, action.Clock[2:4], action.Clock[5:len(action.Clock)-1]),
			action.TeamTricode,
			action.PlayerNameI,
			action.ActionType,
			action.ShotResult,
			fmt.Sprintf(distanceFormat, action.ShotDistance),
			fmt.Sprintf(playByPlayScoreFormat, action.ScoreAway, action.ScoreHome),
		})
	}
	return rows
}

func newPlayByPlayTable(playByPlay *models.PlayByPlayV3) table.Model {
	columns := []table.Column{
		{Title: "Time", Width: 20},
		{Title: "Team", Width: 20},
		{Title: "Player", Width: 20},
		{Title: "Action", Width: 20},
		{Title: "Result", Width: 20},
		{Title: "Dist", Width: 20},
		{Title: "Score", Width: 20},
	}

	rows := buildPlayByPlayRows(playByPlay)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	t.SetStyles(table.Styles{
		Header:   lipgloss.Style{}.Bold(true),
		Cell:     lipgloss.Style{}.Italic(true),
		Selected: selectedItemStyle,
	})
	return t
}

func (m appModel) renderPlayByPlayView() string {
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
		pbpView,
	)
}
