package ui

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

/* TODO: cleanup, put this in a correct file after refactor */

func buildCommonInfoGameHeader(m appModel) string {
	homeTricode := m.choice.homeAbbr
	awayTricode := m.choice.awayAbbr
	homePts := m.choice.homePts
	awayPts := m.choice.awayPts

	away := createBigTeamBadgeStyle(awayTricode).Render(awayTricode)
	home := createBigTeamBadgeStyle(homeTricode).Render(homeTricode)

	var scoreText string
	if awayPts != nil && homePts != nil {
		scoreText = strconv.Itoa(*awayPts) + "  —  " + strconv.Itoa(*homePts)
	} else {
		scoreText = "-  —  -"
	}
	score := boxScoreHeaderScoreStyle.Render(scoreText)

	scoreHeader := boxScoreHeaderRowStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Center, away, "  ", score, "  ", home),
	)

	statusText := preGameStatus
	if homePts != nil && awayPts != nil && *homePts != *awayPts {
		statusText = "FINAL"
	}

	statusHeader := boxScoreHeaderStatusPillStyle.Render(strings.ToUpper(statusText))

	dateHeader := boxScoreHeaderDateStyle.Render(m.date.Format(dateFormat))
	return lipgloss.JoinVertical(lipgloss.Center, scoreHeader, statusHeader, dateHeader)
}
