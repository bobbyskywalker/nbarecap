package ui

import (
	"nbarecap/internal/nba"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 25
	defaultWidth = 120
)

func newGameList() list.Model {
	delegate := gameItemDelegate{}
	l := list.New([]list.Item{}, delegate, defaultWidth, listHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	l.SetShowTitle(false)
	l.Styles.PaginationStyle = paginationStyle
	return l
}

func (m appModel) buildBaseGameInfoList() tea.Cmd {
	return func() tea.Msg {
		games, err := nba.GetAllGamesForDate(&m.date)
		if err != nil {
			return baseGameInfoMsg{nil, err}
		}

		items := make([]list.Item, 0, len(games))
		for _, g := range games {
			items = append(items, gameInfoItem{
				id:       g.GameID,
				awayAbbr: g.Away.Abbr,
				homeAbbr: g.Home.Abbr,
				awayPts:  g.Away.Pts,
				homePts:  g.Home.Pts,
				status:   g.Status,
				arena:    g.Arena,
				tv:       g.NatTV,
			})
		}
		return baseGameInfoMsg{items: items, err: nil}
	}
}

func (m appModel) renderGamesView(header string, footer string) string {
	tableView := listBoxStyle.Render(m.list.View())

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		header+"\n\n"+tableView+"\n\n"+footer,
	)
}
