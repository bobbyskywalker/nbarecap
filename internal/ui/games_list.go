package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listHeight   = 10
	defaultWidth = 20
)

func newGameList() list.Model {
	l := list.New([]list.Item{}, gameItemDelegate{}, defaultWidth, listHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	l.SetShowTitle(false)
	l.Styles.PaginationStyle = paginationStyle
	return l
}

func (m model) buildBaseGameInfoList() tea.Cmd {
	return func() tea.Msg {
		scores, err := m.fetchScoresCmd()
		if err != nil {
			return baseGameInfoMsg{nil, err}
		}

		items := make([]list.Item, 0, len(scores))
		for i, score := range scores {
			if i == 0 {
				continue
			}
			lines := strings.Split(score.GameInfo, "\n")
			items = append(items, gameInfoItem{score.GameId, lines[0]})
		}
		return baseGameInfoMsg{items: items, err: nil}
	}
}
