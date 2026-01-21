package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	goldenYellow    = "#F5C542"
	gray            = "240"
	coolGray        = "#E5E7EB"
	red             = "#C8102E"
	blue            = "#0057B8"
	dotInactiveIcon = "○"
	dotActiveIcon   = "●"
)

var (
	/* List styling */
	gamesListHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(goldenYellow))

	gamesListFooterStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(gray))

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color(coolGray))

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color(red)).
				Bold(true)

	listBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(blue)).
			Bold(true).
			Padding(1, 2)

	paginationStyle = list.DefaultStyles().
			PaginationStyle.
			PaddingLeft(4).
			Foreground(lipgloss.Color("#9CA3AF"))

	/* Table styling */
	awayTeamBadgeStyle = lipgloss.NewStyle().
				Bold(true).
				Padding(0, 1).
				Foreground(lipgloss.Color(coolGray))

	homeTeamBadgeStyle = lipgloss.NewStyle().
				Bold(true).
				Padding(0, 1).
				Background(lipgloss.Color(red)).
				Foreground(lipgloss.Color(coolGray))

	boxScoreHeaderDateStyle = lipgloss.NewStyle().
				Faint(true).
				MarginTop(1)

	dotActiveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(goldenYellow)).
			Margin(0, 1)

	dotInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(gray)).
				Margin(0, 1)
)

func buildBoxScoreHeader(m appModel) string {
	home := awayTeamBadgeStyle.Background(lipgloss.Color(blue)).Render(m.currentBoxScore.HomeTeam.TeamTricode)
	vs := homeTeamBadgeStyle.Render("VS")
	away := awayTeamBadgeStyle.Background(lipgloss.Color(gray)).Render(m.currentBoxScore.AwayTeam.TeamTricode)
	scoreHeader := lipgloss.JoinHorizontal(lipgloss.Center, home, vs, away)
	dateHeader := boxScoreHeaderDateStyle.Render(m.date.Format(dateFormat))
	return lipgloss.JoinVertical(lipgloss.Center, scoreHeader, dateHeader)
}

func buildBoxScoreFooter(m appModel) string {
	var dots string
	if m.showingAway {
		dots = dotInactiveStyle.Render(dotInactiveIcon) + dotActiveStyle.Render(dotActiveIcon)
	} else {
		dots = dotActiveStyle.Render(dotActiveIcon) + dotInactiveStyle.Render(dotInactiveIcon)
	}
	return lipgloss.JoinVertical(lipgloss.Center, dots)
}
