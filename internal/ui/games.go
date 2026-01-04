package ui

import (
	"log"
	"nbarecap/internal/recaps"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	date       time.Time
	gameScores []string
	err        error
	style      *style

	viewport   viewport.Model
	fetchReady bool
	termWidth  int
	termHeight int
}

type style struct {
	border      lipgloss.Border
	borderColor lipgloss.Color
}

type scoresMsg struct {
	gameScores []string
	err        error
}

func newGameBoxStyle() *style {
	return &style{
		border:      lipgloss.RoundedBorder(),
		borderColor: "36",
	}
}

func (s *style) Top() lipgloss.Style {
	/* TODO: date intro + cool msg */
	return lipgloss.NewStyle().
		Border(s.border).
		BorderForeground(s.borderColor)
}

/* TODO: prettier boxes for games :) */
func (s *style) Box() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(s.border).
		BorderForeground(s.borderColor)
}

func renderBoxes(m model) string {
	box := m.style.Box()

	var boxes []string
	for _, score := range m.gameScores {
		boxes = append(boxes, box.Render(score))
	}
	return lipgloss.JoinVertical(lipgloss.Center, boxes...)
}

func fetchScoresCmd(date time.Time) tea.Cmd {
	return func() tea.Msg {
		scores, err := recaps.GetAllGamesForDate(&date)
		return scoresMsg{scores, err}
	}
}

func (m model) Init() tea.Cmd {
	return fetchScoresCmd(m.date)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.termWidth, m.termHeight = msg.Width, msg.Height
		headerHeight := 2
		footerHeight := 1

		if !m.fetchReady {
			m.viewport = viewport.New(msg.Width, msg.Height-headerHeight-footerHeight)
			m.viewport.YPosition = headerHeight
			m.fetchReady = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - headerHeight - footerHeight
		}

		m.viewport.SetContent(renderBoxes(m))
		return m, nil

	case scoresMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
		m.gameScores = msg.gameScores

		if m.fetchReady {
			m.viewport.SetContent(renderBoxes(m))
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			m.date = m.date.AddDate(0, 0, -1)
			return m, fetchScoresCmd(m.date)
		case "right":
			m.date = m.date.AddDate(0, 0, 1)
			return m, fetchScoresCmd(m.date)
		case "up", "k":
			m.viewport.LineUp(1)
			return m, nil
		case "down", "j":
			m.viewport.LineDown(1)
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}
	if !m.fetchReady {
		return "Loading..."
	}

	header := lipgloss.NewStyle().
		Bold(true).
		Render(m.date.Format("2006-01-02") + "  ←/→ change day • ↑/↓ scroll • q quit")

	footer := lipgloss.NewStyle().
		Bold(true).
		Render("Press q to quit")

	return header + "\n" + m.viewport.View() + "\n" + footer
}

func RunGamesView(date time.Time) {
	f, err := tea.LogToFile("tea.log", "debug")
	if err != nil {
		log.Fatalf("Error: %w", err)
	}
	f.Close()

	m := model{date: date, style: newGameBoxStyle()}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
