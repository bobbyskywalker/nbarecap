package ui

import (
	"nbarecap/internal/utils"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	goldenYellow    = "#F5C542"
	gray            = "240"
	coolGray        = "#E5E7EB"
	red             = "#C8102E"
	blue            = "#0057B8"
	black           = "#000000"
	dotInactiveIcon = "○"
	dotActiveIcon   = "●"
)

var TeamColors = map[string]string{
	"ATL": "196",
	"BOS": "34",
	"BKN": "245",
	"CHA": "62",
	"CHI": "196",
	"CLE": "52",
	"DAL": "25",
	"DEN": "178",
	"DET": "124",
	"GSW": "227",
	"HOU": "160",
	"IND": "21",
	"LAC": "12",
	"LAL": "93",
	"MEM": "69",
	"MIA": "161",
	"MIL": "22",
	"MIN": "27",
	"NOP": "17",
	"NYK": "208",
	"OKC": "39",
	"ORL": "27",
	"PHI": "21",
	"PHX": "129",
	"POR": "160",
	"SAC": "99",
	"SAS": "250",
	"TOR": "124",
	"UTA": "60",
	"WAS": "124",
}

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
	homeTeamBadgeStyle = lipgloss.NewStyle().
				Bold(true).
				Padding(0, 1).
				Background(lipgloss.Color(coolGray)).
				Foreground(lipgloss.Color(black))

	/* boxScore header styling */
	boxScoreHeaderRowStyle = lipgloss.NewStyle().
				MarginTop(1).
				MarginBottom(1)

	boxScoreHeaderScoreStyle = lipgloss.NewStyle().
					Bold(true).
					Padding(1, 4).
					Border(lipgloss.DoubleBorder()).
					BorderForeground(lipgloss.Color(goldenYellow)).
					Foreground(lipgloss.Color(goldenYellow)).
					Align(lipgloss.Center)

	boxScoreHeaderStatusPillStyle = lipgloss.NewStyle().
					Bold(true).
					Padding(0, 2).
					Border(lipgloss.RoundedBorder()).
					BorderForeground(lipgloss.Color(gray)).
					Foreground(lipgloss.Color(coolGray)).
					Align(lipgloss.Center)

	boxScoreHeaderDateStyle = lipgloss.NewStyle().
				Faint(true).
				MarginTop(1)

	/* controls styling */
	dotActiveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(goldenYellow)).
			Margin(0, 1)

	dotInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(gray)).
				Margin(0, 1)
)

func createTeamBadgeStyle(tricode string) lipgloss.Style {
	bg := TeamColors[tricode]
	bgNum, _ := strconv.Atoi(bg)

	fg := "15"
	if utils.IsLightANSI(bgNum) {
		fg = "0"
	}

	return lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		Background(lipgloss.Color(bg)).
		Foreground(lipgloss.Color(fg))
}

func createBigTeamBadgeStyle(tricode string) lipgloss.Style {
	return createTeamBadgeStyle(tricode).
		Padding(1, 4).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("0")).
		Align(lipgloss.Center)
}
