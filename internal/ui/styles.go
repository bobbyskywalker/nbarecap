package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

/* List styling */
var (
	gamesListHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#F5C542"))

	gamesListFooterStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("240"))

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("#E5E7EB")) // cool gray

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("#C8102E")). // red
				Bold(true)

	paginationStyle = list.DefaultStyles().
			PaginationStyle.
			PaddingLeft(4).
			Foreground(lipgloss.Color("#9CA3AF"))

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4).
			Foreground(lipgloss.Color("#F9FAFB"))

	tableBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#0057B8")).
			Bold(true).
			Padding(1, 2)
)
