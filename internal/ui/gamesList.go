package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func newGameList(date time.Time) list.Model {
	l := list.New([]list.Item{}, gameItemDelegate{}, defaultWidth, listHeight)
	l.Title = fmt.Sprintf(titleFormat, date.Format(dateFormat))
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
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
			lines := strings.Split(score, "\n")
			items = append(items, gameInfoItem(lines[0]))
		}
		return baseGameInfoMsg{items: items, err: nil}
	}
}
