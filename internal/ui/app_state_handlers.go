package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func updateDate(date time.Time, dateDelta int) time.Time {
	return date.AddDate(0, 0, dateDelta)
}

func updateGamesState(m appModel, msg tea.Msg) (appModel, tea.Cmd) {
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
}
