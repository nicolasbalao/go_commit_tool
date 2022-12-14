package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nicolasbalao/go_commit_toll/style"
)

var componentStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("63")).
	AlignHorizontal(lipgloss.Center).Padding(0, 1)

var selectedStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("63")).
	Foreground(lipgloss.Color("230"))
var unselectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("60"))

type confirmModel struct {
	title    string
	question string
	choice   bool
}

func newConfirmComponent(title string, question string) *confirmModel {
	return &confirmModel{
		title:    title,
		question: question,
		choice:   false,
	}
}

func Init() tea.Cmd {
	return nil
}

func (m *confirmModel) Update(msg tea.Msg, tm Model) (bool, tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "h", "left":
			m.choice = true
		case "right", "l":
			m.choice = false
		case "enter":
			if m.choice {
				tm.state++
				return m.choice, tm, nil
			}
			tm.state += 2
			if tm.state > 9 {
				tm.state--
			}
			return m.choice, tm, nil
		}
	}
	return m.choice, tm, nil
}

func (m confirmModel) View() string {
	if m.choice {
		return style.TitleStyle.Render(m.title) + "\n" + componentStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.question,
				"",
				lipgloss.JoinHorizontal(lipgloss.Left, selectedStyle.Render("yes"), " ", "no"),
			),
		)

	}
	return style.TitleStyle.Render(m.title) + "\n" + componentStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.question,
			"",
			lipgloss.JoinHorizontal(lipgloss.Left, "yes", " ", selectedStyle.Render("no")),
		),
	)
}
