package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	scoreFormat     = "%3d - %-3d"
	emptyScoreStr   = "   @   "
	cardWidth       = 50
	lineWidth       = 40
	gameEndedStatus = "final"
	halftimeStatus  = "half"
	quarterIdAlias  = "Q"
	arenaEmoji      = "🏟 "
)

type gameItemDelegate struct{}

func (d gameItemDelegate) Height() int                             { return 2 }
func (d gameItemDelegate) Spacing() int                            { return 1 }
func (d gameItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d gameItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	game, ok := listItem.(gameInfoItem)
	if !ok {
		return
	}

	selected := index == m.Index()

	/* badge render */
	awayBadge := createTeamBadgeStyle(game.awayAbbr).Render(game.awayAbbr)
	homeBadge := createTeamBadgeStyle(game.homeAbbr).Render(game.homeAbbr)

	/* score render */
	var scoreStr string
	if game.awayPts != nil && game.homePts != nil {
		scoreStr = fmt.Sprintf(scoreFormat, *game.awayPts, *game.homePts)
	} else {
		scoreStr = emptyScoreStr
	}
	score := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(coolGray)).Render(scoreStr)

	teamsAndScoreLine := lipgloss.JoinHorizontal(lipgloss.Center, awayBadge, " ", score, " ", homeBadge)

	/* status style build */
	statusText := strings.TrimSpace(game.status)
	var statusStyle lipgloss.Style
	switch {
	case strings.Contains(strings.ToLower(statusText), gameEndedStatus):
		statusStyle = gameEndedStatusStyle
	case strings.Contains(statusText, quarterIdAlias) || strings.Contains(strings.ToLower(statusText), halftimeStatus):
		statusStyle = gameInProgressStatusStyle
	default:
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(coolGray))
	}
	status := statusStyle.Render(statusText)

	/* status and arena line render */
	separator := statusAndArenaSeparatorStyle
	statusAndArenaLine := status
	if arena := strings.TrimSpace(game.arena); arena != "" {
		arenaStyled := arenaStyle.Render(arenaEmoji + arena)
		statusAndArenaLine = status + separator + arenaStyled
	}

	teamsAndScoreLineStyle := lipgloss.NewStyle().Width(lineWidth).Align(lipgloss.Center)
	statusAndArenaLineStyle := lipgloss.NewStyle().Width(lineWidth).Align(lipgloss.Center)

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		teamsAndScoreLineStyle.Render(teamsAndScoreLine),
		statusAndArenaLineStyle.Render(statusAndArenaLine),
	)

	itemStyle := cardStyle

	if selected {
		itemStyle = itemStyle.
			Border(lipgloss.Border{Left: "┃"}, false, false, false, true).
			BorderForeground(lipgloss.Color(goldenYellow)).
			Background(lipgloss.Color("#1F1F1F"))
	}

	_, _ = fmt.Fprint(w, itemStyle.Render(content))
}
