package ui

import (
	"fmt"
	"log"
	"nbarecap/internal/recaps"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 15
	defaultWidth = 20
	dateFormat   = "2006-01-02"
	titleFormat  = "Games on %s"
	dayStep      = 1
)

/* Models */
type model struct {
	date       time.Time
	gameScores []string
	numGames   int
	err        error

	list   list.Model
	choice string

	termWidth  int
	termHeight int
}

type gameInfoItem string

func (g gameInfoItem) FilterValue() string { return "" }

/* Tea messages */
type baseGameInfoMsg struct {
	items []list.Item
	err   error
}

/* Games API */
func (m model) fetchScoresCmd() ([]string, error) {
	scores, err := recaps.GetAllGamesForDate(&m.date)
	m.gameScores = scores
	return scores, err
}

func updateDates(date time.Time, dateDelta int) (time.Time, string) {
	date = date.AddDate(0, 0, dateDelta)
	title := fmt.Sprintf(titleFormat, date.Format(dateFormat))
	return date, title
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
			m.date, m.list.Title = updateDates(m.date, -dayStep)
			return m, m.buildBaseGameInfoList()

		case "right":
			m.date, m.list.Title = updateDates(m.date, dayStep)
			return m, m.buildBaseGameInfoList()

		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(gameInfoItem)
			if ok {
				m.choice = string(i)
			}
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
		Foreground(lipgloss.Color("240")).
		Render("←/→ change day • ↑/↓ scroll • q quit")

	footer := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("240")).
		Render(fmt.Sprintf("Showing %d games", m.numGames))

	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("Selected game %s!", m.choice))
	}

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		header+"\n"+m.list.View()+"\n"+footer,
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
		list: newGameList(date),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
