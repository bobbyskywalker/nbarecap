package ui

import (
	"log"
	"nbarecap/internal/recaps"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Date   string
	Scores string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.Scores
}

func (m Model) RunGamesView() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("Error: %w", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	/* tmp */
	m.Scores, _ = recaps.GetAllGamesForDate(&m.Date)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
