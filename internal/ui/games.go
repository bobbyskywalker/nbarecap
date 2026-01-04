package ui

import (
	"log"
	"nbarecap/internal/recaps"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	date       time.Time
	gameScores []string
	style      *style
	termWidth  int
	termHeight int
	err        error
}

type style struct {
	border      lipgloss.Border
	borderColor lipgloss.Color
}

type scoresMsg struct {
	gameScores []string
	err        error
}

func gameBoxStyle() *style {
	return &style{
		border:      lipgloss.RoundedBorder(),
		borderColor: "36",
	}
}

/* TODO: some clever, fitting on screen box style */
func (s *style) Box() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(s.border).
		BorderForeground(s.borderColor)
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
		return m, nil

	case scoresMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
		m.gameScores = msg.gameScores

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
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}
	if len(m.gameScores) == 0 {
		return "Loading games..."
	}

	box := m.style.Box()

	var boxes []string
	for _, score := range m.gameScores {
		boxes = append(boxes, box.Render(score))
	}

	column := lipgloss.JoinVertical(lipgloss.Center, boxes...)
	return lipgloss.NewStyle().Align(lipgloss.Center).Render(column)
}

func RunGamesView(date time.Time) {
	f, err := tea.LogToFile("tea.log", "debug")
	if err != nil {
		log.Fatalf("Error: %w", err)
	}
	f.Close()

	m := model{date: date, style: gameBoxStyle()}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
