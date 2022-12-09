package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type bodyModel struct {
	focusIndex int
	inputs     []textinput.Model
}

func newBodyModel() *bodyModel {
	m := bodyModel{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Prompt = "Scope: "
			t.Placeholder = " api"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Prompt = "Description: "
			t.Placeholder = "change the endpoint"
		}

		m.inputs[i] = t
	}

	return &m
}

func (m *bodyModel) Update(msg tea.Msg, tm Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		// set focus to next input
		case "tab", "ctrl+p", "enter", "down", "ctrl+n", "up":

			if keypress == "enter" && m.focusIndex == len(m.inputs) {
				tm.state++
				return tm, nil
			}

			// change inputs focus
			if keypress == "up" || keypress == "ctrl+p" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > 4 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = 4
			}

			// cmds := make([]tea.Cmd, len(m.inputs))

			// Set focus
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					m.inputs[i].Focus()
					m.inputs[i].TextStyle = focusedStyle
					m.inputs[i].PromptStyle = focusedStyle
					continue
				}
				m.inputs[i].Blur()
				m.inputs[i].TextStyle = noStyle
				m.inputs[i].PromptStyle = noStyle
			}

			tm.commit.scope = m.inputs[0].Value()
			tm.commit.description = m.inputs[1].Value()

			return tm, nil
		}
	}
	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	return tm, cmd
}

func (m *bodyModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

func (m bodyModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())

		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
