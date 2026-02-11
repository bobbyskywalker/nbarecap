package games

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func updateViewSelectionState(m appModel, msg tea.Msg) (appModel, tea.Cmd) {
	switch msg := msg.(type) {
	case viewSelectionMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.optionsList.SetItems(msg.items)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "backspace":
			m.state = games
			return m, nil
		case "enter":
			selected := m.optionsList.SelectedItem()
			if selected == nil {
				return m, nil
			}
			switch selected.(viewSelectionItem).value {
			case "Box Score":
				m.showingAway = true
				m.showingPercentages = false
				m.state = boxScore
				m.err = nil
				m.currentBoxScore = nil
				return m, fetchBoxScoreCmd(m.choice.id)
			case "Play By Play":
				m.state = playByPlay
				m.err = nil
				m.selectedPeriod = 1
				m.currentPlayByPlay = nil
				return m, m.fetchPlayByPlayCmd(m.choice.id)
			}
		}
	}

	var cmd tea.Cmd
	m.optionsList, cmd = m.optionsList.Update(msg)
	return m, cmd
}

func updateDate(date time.Time, dateDelta int) time.Time {
	return date.AddDate(0, 0, dateDelta)
}

func updateGamesState(m appModel, msg tea.Msg) (appModel, tea.Cmd) {
	var cmd tea.Cmd
	m.gamesList, cmd = m.gamesList.Update(msg)

	switch msg := msg.(type) {
	case baseGameInfoMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.err = nil
		m.gamesList.SetItems(msg.items)
		m.numGames = len(msg.items)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			m.date = updateDate(m.date, -dayStep)
			return m, gamesListCmd(&m.date)

		case "right":
			m.date = updateDate(m.date, dayStep)
			return m, gamesListCmd(&m.date)

		case "enter":
			it, ok := m.gamesList.SelectedItem().(gameInfoItem)
			if !ok {
				return m, cmd
			}
			m.choice = &it
			m.state = viewSelection
			return m, buildViewSelectionMenu()
		}
	}
	return m, cmd
}

func updateBoxScoreState(m appModel, msg tea.Msg) (appModel, tea.Cmd) {
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
		m.boxTable = newBoxScoreTable(msg.score, m.showingAway, m.showingPercentages)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "backspace":
			m.state = games
			return m, nil
		case "left", "right":
			m.showingAway = !m.showingAway
			if m.currentBoxScore != nil {
				m.boxTable = newBoxScoreTable(m.currentBoxScore, m.showingAway, m.showingPercentages)
			}
			return m, nil
		case "m":
			m.showingPercentages = !m.showingPercentages
			if m.currentBoxScore != nil {
				m.boxTable = newBoxScoreTable(m.currentBoxScore, m.showingAway, m.showingPercentages)
			}
			return m, nil
		}
	}
	return m, cmd
}

func updatePeriod(period int, delta int, maxPeriod int) int {
	period += delta
	if period < 1 {
		period = 1
	} else if period > maxPeriod {
		period = maxPeriod
	}
	return period
}

func updatePlayByPlayState(m appModel, msg tea.Msg) (appModel, tea.Cmd) {
	var cmd tea.Cmd
	m.playByPlayTable, cmd = m.playByPlayTable.Update(msg)

	switch msg := msg.(type) {
	case playByPlayMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.err = nil
		m.currentPlayByPlay = msg.content
		m.maxPeriod = getMaxPeriod(msg.content)
		m.playByPlayTable = newPlayByPlayTable(msg.content, m.selectedPeriod)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "backspace":
			m.state = games
			return m, nil
		case "left":
			m.selectedPeriod = updatePeriod(m.selectedPeriod, -1, m.maxPeriod)
			m.playByPlayTable = newPlayByPlayTable(m.currentPlayByPlay, m.selectedPeriod)
			return m, nil
		case "right":
			m.selectedPeriod = updatePeriod(m.selectedPeriod, 1, m.maxPeriod)
			m.playByPlayTable = newPlayByPlayTable(m.currentPlayByPlay, m.selectedPeriod)
			return m, nil
		}
		return m, cmd
	}
	return m, cmd
}
