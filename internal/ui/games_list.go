package ui

import (
	"nbarecap/internal/nba"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	gamesListDefaultHeight = 25
	gamesListDefaultWidth  = 120
)

type baseGameInfoMsg struct {
	items []list.Item
	err   error
}

type gameInfoItem struct {
	id       string
	awayAbbr string
	homeAbbr string
	awayPts  *int
	homePts  *int
	status   string
	arena    string
	tv       string
}

func (g gameInfoItem) FilterValue() string { return "" }

func newGameList() list.Model {
	delegate := gameItemDelegate{}
	l := list.New([]list.Item{}, delegate, gamesListDefaultWidth, gamesListDefaultHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	l.SetShowTitle(false)
	l.Styles.PaginationStyle = paginationStyle
	return l
}

func gamesListCmd(date *time.Time) tea.Cmd {
	return func() tea.Msg {
		games, err := nba.GetAllGamesForDate(date)
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
	tableView := listBoxStyle.Render(m.gamesList.View())

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		header+"\n\n"+tableView+"\n\n"+footer,
	)
}
