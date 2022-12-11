package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nicolasbalao/go_commit_toll/style"
)

type textAreaModel struct {
	title    string
	textarea textarea.Model
	err      error
}

func newTexteAreaComponent(title string, placeholder string) *textAreaModel {
	ta := textarea.New()
	ta.Placeholder = placeholder
	ta.Focus()


	return &textAreaModel{
        title: title,
		textarea: ta,
		err:      nil,
	}
}

func (m textAreaModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m *textAreaModel) Update(msg tea.Msg, tm Model) (string, tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

    tm.focusedTextArea = m.textarea.Focused()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case "enter":
			if !m.textarea.Focused() {
				tm.state++
				valueTextArea := m.textarea.Value()
				return valueTextArea, tm, nil
			}

		}
	case errMsg:
		m.err = msg
		return "", tm, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return "", tm, tea.Batch(cmds...)
}

func (m textAreaModel) View() string {
	return fmt.Sprintf(
		"%s \n\n%s\n\n",
        style.TitleStyle.Render(m.title),
		m.textarea.View(),
	) + "\n\n"
}
