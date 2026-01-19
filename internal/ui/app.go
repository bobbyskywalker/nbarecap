package ui

import (
	"fmt"
	"log"
	"nbarecap/internal/nba"
	"nbarecap/pkg/nba_api/models"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dateFormat         = "2006-01-02"
	dayStep            = 1
	boxScoreLoadingMsg = "\nLoading box score...\n\n(esc to go back)"
	boxScoreLoadedMsg  = "\n\n(esc to go back)"
)

type appState int

const (
	games appState = iota
	boxScore
)

/* Models */
type model struct {
	/* Games list data */
	date       time.Time
	gameScores []models.GameInfoFormatted
	numGames   int
	list       list.Model
	choice     *gameInfoItem

	/* Box score data */
	boxTable        table.Model
	currentBoxScore *models.BoxScoreTraditionalV3

	/* Viewport params */
	termWidth  int
	termHeight int

	/* State */
	state appState
	err   error
}

/* List items */
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

type boxScoreMsg struct {
	score *models.BoxScoreTraditionalV3
	err   error
}

/* NBA API */
func (m model) fetchScoresCmd() ([]models.GameInfoFormatted, error) {
	scores, err := nba.GetAllGamesForDate(&m.date)
	if err != nil {
		return nil, err
	}
	m.gameScores = scores
	return scores, err
}

func (m model) fetchBoxScoreCmd(gameID string) tea.Cmd {
	return func() tea.Msg {
		bx, err := nba.GetBoxScoreForGame(gameID)
		if err != nil {
			return boxScoreMsg{err: err}
		}
		return boxScoreMsg{score: bx}
	}
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

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch m.state {

	case games:
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)

		switch msg := msg.(type) {
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

			case "enter":
				it, ok := m.list.SelectedItem().(gameInfoItem)
				if !ok {
					return m, cmd
				}
				m.choice = &it
				m.state = boxScore
				m.err = nil
				m.currentBoxScore = nil

				return m, tea.Batch(cmd, m.fetchBoxScoreCmd(it.id))
			}
		}

		return m, cmd

	case boxScore:
		var cmd tea.Cmd
		m.boxTable, cmd = m.boxTable.Update(msg)

		switch msg := msg.(type) {
		case boxScoreMsg:
			if msg.err != nil {
				m.err = msg.err
				return m, nil
			}
			m.err = nil
			m.currentBoxScore = msg.score
			m.boxTable = newBoxScoreTable(msg.score)
			return m, nil

		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "backspace":
				m.state = games
				return m, nil
			}
		}

		return m, cmd

	default:
		panic("unknown state")
	}
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

	switch m.state {

	case games:
		tableView := tableBoxStyle.Render(m.list.View())

		return lipgloss.Place(
			m.termWidth,
			m.termHeight,
			lipgloss.Center,
			lipgloss.Center,
			header+"\n"+tableView+"\n"+footer,
		)

	case boxScore:
		if m.currentBoxScore == nil {
			return lipgloss.Place(
				m.termWidth,
				m.termHeight,
				lipgloss.Center,
				lipgloss.Center,
				header+boxScoreLoadingMsg,
			)
		}

		boxView := tableBoxStyle.Render(m.boxTable.View())

		return lipgloss.Place(
			m.termWidth,
			m.termHeight,
			lipgloss.Center,
			lipgloss.Center,
			header+"\n"+boxView+boxScoreLoadedMsg,
		)

	default:
		return ""
	}
}

func StartUi(date time.Time) {
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
