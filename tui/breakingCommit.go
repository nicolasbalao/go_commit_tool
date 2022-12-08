package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type breakingModel struct {
	textInput textinput.Model
	err       error
}

func newBreakingModel() *breakingModel {
	ti := textinput.New()
	ti.Placeholder = "(y/n)"
	ti.Focus()
	ti.CharLimit = 1
    ti.Width = 20

	return &breakingModel{
		textInput: ti,
		err:       nil,
	}
}

func breakingUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case errMsg:
		m.breakingComponent.err = msg
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			// set breaking change to false or true depend on the answer
			if inputValue := m.breakingComponent.textInput.Value(); inputValue == "n" {
				m.commit.breaking = false
			} else {
				m.commit.breaking = true
			}
			m.state++
			return m, nil
		}
	}

	// need ?
	m.breakingComponent.textInput, cmd = m.breakingComponent.textInput.Update(msg)

	return m, cmd
}

func (m breakingModel) View() string {
	return fmt.Sprintf("Commit have breaking change? %s \n", m.textInput.View())
}
