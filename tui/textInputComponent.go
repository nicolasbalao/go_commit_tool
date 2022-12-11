package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nicolasbalao/go_commit_toll/style"
)

type textInputModel struct {
	title     string
	textInput textinput.Model
	err       error
}

func newTextInputComponent(title string, placeholder string) *textInputModel {

	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.PromptStyle = style.InputPrompStyle


	return &textInputModel{
		title:     title,
		textInput: ti,
	}
}

func (m textInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *textInputModel) Update(msg tea.Msg, tm Model) (string, tea.Model, tea.Cmd) {
	var cmd tea.Cmd


	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			tm.state++
		}
	case errMsg:
		m.err = msg
	}

	m.textInput, cmd = m.textInput.Update(msg)
	value := m.textInput.Value()
	return value, tm, cmd
}

func (m textInputModel) View() string {
	return fmt.Sprintf(
		("%s\n\n%s\n\n"),
		style.TitleStyle.Render(m.title),
		style.InputTextStyle.Render(m.textInput.View()),
	) + "\n"
}
