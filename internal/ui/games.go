package ui

import (
	"log"
	"nbarecap/internal/recaps"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Date       time.Time
	gameScores []string
	err        error
}

type scoresMsg struct {
	gameScores []string
	err        error
}

func fetchScoresCmd(date time.Time) tea.Cmd {
	return func() tea.Msg {
		scores, err := recaps.GetAllGamesForDate(&date)
		return scoresMsg{scores, err}
	}
}

func (m Model) Init() tea.Cmd {
	return fetchScoresCmd(m.Date)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

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
			m.Date = m.Date.AddDate(0, 0, -1)
			return m, fetchScoresCmd(m.Date)
		case "right":
			m.Date = m.Date.AddDate(0, 0, 1)
			return m, fetchScoresCmd(m.Date)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.gameScores
}

func RunGamesView(date time.Time) {
	f, err := tea.LogToFile("tea.log", "debug")
	if err != nil {
		log.Fatalf("Error: %w", err)
	}
	f.Close()

	m := Model{Date: date}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
