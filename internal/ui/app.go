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
)

const (
	dateFormat               = "2006-01-02"
	dayStep                  = 1
	boxScoreLoadingMsg       = "\nLoading box score...\n\n(esc to go back)"
	gamesListHeaderMsgFormat = "NBARECAP • %s"
	gamesListFooterMsgFormat = "Showing %d games"
)

type appState int

const (
	games appState = iota
	boxScore
)

/* Models */
type appModel struct {
	/* Games list data */
	date       time.Time
	gameScores []models.GameInfoFormatted
	numGames   int
	list       list.Model
	choice     *gameInfoItem

	/* Box score data */
	boxTable        table.Model
	currentBoxScore *models.BoxScoreTraditionalV3
	showingAway     bool

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
func (m appModel) fetchScoresCmd() ([]models.GameInfoFormatted, error) {
	scores, err := nba.GetAllGamesForDate(&m.date)
	if err != nil {
		return nil, err
	}
	m.gameScores = scores
	return scores, err
}

func (m appModel) fetchBoxScoreCmd(gameID string) tea.Cmd {
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

func (m appModel) Init() tea.Cmd {
	return m.buildBaseGameInfoList()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.showingAway = true
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
			m.boxTable = newBoxScoreTable(msg.score, m.showingAway)
			return m, nil

		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "backspace":
				m.state = games
				return m, nil
			case "left", "right":
				m.showingAway = !m.showingAway
				if m.currentBoxScore != nil {
					m.boxTable = newBoxScoreTable(m.currentBoxScore, m.showingAway)
				}
				return m, nil
			}
		}

		return m, cmd

	default:
		panic("unknown state")
	}
}

func (m appModel) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	var header string
	var footer string

	switch m.state {
	case games:
		header = gamesListHeaderStyle.Render(fmt.Sprintf(gamesListHeaderMsgFormat, m.date.Format(dateFormat)))
		footer = gamesListFooterStyle.Render(fmt.Sprintf(gamesListFooterMsgFormat, m.numGames))
		return m.renderGamesView(header, footer)
	case boxScore:
		var header string
		if m.currentBoxScore != nil {
			header = buildBoxScoreHeader(m)
		}
		footer = buildBoxScoreFooter(m)
		return m.renderBoxScoreView(header, footer)
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

	m := appModel{
		date: date,
		list: newGameList(),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
