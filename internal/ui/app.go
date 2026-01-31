package ui

import (
	"fmt"
	"log"
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
	viewSelection
	boxScore
	playByPlay
)

/* Models */
type appModel struct {
	/* Games gamesList data */
	date      time.Time
	numGames  int
	gamesList list.Model
	choice    *gameInfoItem

	/* viewSelection data */
	optionsList  list.Model
	selectedView string

	/* Box score data */
	boxTable           table.Model
	currentBoxScore    *models.BoxScoreTraditionalV3
	showingAway        bool
	showingPercentages bool

	/* Viewport params */
	termWidth  int
	termHeight int

	/* State */
	state appState
	err   error
}

func (m appModel) Init() tea.Cmd {
	return m.buildBaseGameInfoList()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.termWidth, m.termHeight = msg.Width, msg.Height
		m.gamesList.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch m.state {

	case games:
		return updateGamesState(m, msg)

	case boxScore:
		return updateBoxScoreState(m, msg)

	case viewSelection:
		return updateViewSelectionState(m, msg)

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
	case viewSelection:
		return m.renderViewSelectionMenu()
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
		date:        date,
		gamesList:   newGameList(),
		optionsList: newViewSelectionMenu(),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
