package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

/* List styling */
var (
	titleStyle = lipgloss.NewStyle().
			MarginLeft(2).
			Bold(true).
			Underline(true).
			Foreground(lipgloss.Color("#0057B8")) // blue-ish

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

	helpStyle = list.DefaultStyles().
			HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1).
			Foreground(lipgloss.Color("#9CA3AF"))

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4).
			Foreground(lipgloss.Color("#F9FAFB"))
)
