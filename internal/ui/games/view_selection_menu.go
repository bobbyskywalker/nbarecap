package games

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	viewSelectionDefaultWidth  = 20
	viewSelectionDefaultHeight = 10
)

type viewSelectionMsg struct {
	items []list.Item
	err   error
}

type viewSelectionItem struct {
	value string
}

func (v viewSelectionItem) FilterValue() string {
	return ""
}

func (v viewSelectionItem) Title() string {
	return v.value
}

func (v viewSelectionItem) Description() string {
	return ""
}

func newViewSelectionMenu() list.Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), viewSelectionDefaultWidth, viewSelectionDefaultHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)
	return l
}

func buildViewSelectionMenu() tea.Cmd {
	return func() tea.Msg {
		var items []list.Item
		items = append(items, viewSelectionItem{value: "Box Score"})
		items = append(items, viewSelectionItem{value: "Play By Play"})
		return viewSelectionMsg{
			items: items,
			err:   nil,
		}
	}
}

func renderViewSelectionMenu(m appModel) string {
	tableView := listBoxStyle.Render(m.optionsList.View())

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		tableView,
	)
}
