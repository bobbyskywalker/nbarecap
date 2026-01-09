package ui

import (
	"fmt"
	"log"
	"nbarecap/internal/models"
	"nbarecap/internal/nba"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dateFormat = "2006-01-02"
	dayStep    = 1
)

/* Models */
type model struct {
	date       time.Time
	gameScores []models.GameInfoFormatted
	numGames   int
	err        error

	list   list.Model
	choice *gameInfoItem

	termWidth  int
	termHeight int
}

type gameInfoItem struct {
	id    string
	value string
}

func (g gameInfoItem) FilterValue() string { return "" }

/* Tea messages */
type baseGameInfoMsg struct {
	items []list.Item
	err   error
}

/* Games API */
func (m model) fetchScoresCmd() ([]models.GameInfoFormatted, error) {
	scores, err := nba.GetAllGamesForDate(&m.date)
	if err != nil {
		return nil, err
	}
	m.gameScores = scores
	return scores, err
}

func updateDate(date time.Time, dateDelta int) time.Time {
	return date.AddDate(0, 0, dateDelta)
}

/* Tea program */

func (m model) Init() tea.Cmd {
	return m.buildBaseGameInfoList()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.termWidth, m.termHeight = msg.Width, msg.Height
		m.list.SetWidth(msg.Width)
		return m, nil

	case baseGameInfoMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.err = nil
		m.list.SetItems(msg.items)
		m.numGames = len(msg.items)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			m.date = updateDate(m.date, -dayStep)
			return m, m.buildBaseGameInfoList()

		case "right":
			m.date = updateDate(m.date, dayStep)
			return m, m.buildBaseGameInfoList()

		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(gameInfoItem)
			if ok {
				m.choice = &i
			}
			nba.FetchBoxScoreForGame(i.id)
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#F5C542")).
		Render(fmt.Sprintf("NBARECAP • %s", m.date.Format(dateFormat)))

	footer := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("240")).
		Render(fmt.Sprintf("Showing %d games", m.numGames))

	if m.choice != nil {
		return quitTextStyle.Render(fmt.Sprintf("Selected game %s!", m.choice))
	}

	table := tableBoxStyle.Render(m.list.View())

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		header+"\n"+table+"\n"+footer,
	)
}

func RunGamesView(date time.Time) {
	f, err := tea.LogToFile("tea.log", "debug")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer f.Close()

	m := model{
		date: date,
		list: newGameList(),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
