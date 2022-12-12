package tui

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nicolasbalao/go_commit_toll/style"
)

// Structures and enums

type commitMessage struct {
	typeCommit   string
	scope        string // optional
	description  string
	body         string // optional
	breaking     bool
	breakingDesc string
	footer       string // optional
}

// Enums for the state of the app
type State int

const (
	typeS State = iota
	breakingS
	breakingDescS
	scopeS
	descriptionS
	bodyS
	footerS
	previewS
	commitS
)

// Global struct of the app

type Model struct {
	typeComponent         *typeCommitModel
	breakingComponent     *confirmModel
	breakingDescComponent *textInputModel
	scopeComponent        *textInputModel
	descriptionComponent  *textInputModel
	bodyComponent         *textAreaModel
	footerComponent       *textAreaModel
	previewComponent      *confirmModel

	focusedTextArea bool
	commit          *commitMessage
	state           State
}

// Create the Model
func NewModel() Model {

	commitMessage := commitMessage{
		scope:  "", //optional
		body:   "", //optional
		footer: "", //optional
	}

	return Model{
		typeComponent:     newTypeModel(),
		breakingComponent: newConfirmComponent("breaking change", "Have breaking change ?"),
		breakingDescComponent: newTextInputComponent(
			"description of the breaking change",
			"short desc",
		),
		scopeComponent:       newTextInputComponent("scope", "api"),
		descriptionComponent: newTextInputComponent("description", "short description"),
		bodyComponent: newTexteAreaComponent(
			"Body of the commit",
			"long description of the commit",
		),
		footerComponent: newTexteAreaComponent(
			"footer",
			"description of the breaking change and ref if you want",
		),

		previewComponent: newConfirmComponent("Preview", "Commit ?"),

		commit: &commitMessage,
		state:  typeS,
	}
}

// Init

func (m Model) Init() tea.Cmd {
	return nil
}

// Update

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// switch on the type of msg
	switch msg := msg.(type) {
	case tea.KeyMsg:
		//switch on the value of the msg
		switch keypress := msg.String(); keypress {
		// esc or ctrl+c for quit the app
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if !m.focusedTextArea {
				return m, tea.Quit
			}
		}

	}
	// Switch on the state of the app
	switch m.state {
	//Call update function of the components
	case typeS:
		value, rm, cmd := m.typeComponent.Update(msg, m)
        m.commit.typeCommit = value
		return rm, cmd
	case breakingS:
		value, rm, cmd := m.breakingComponent.Update(msg, m)
		m.commit.breaking = value
		return rm, cmd
	case breakingDescS:
		value, rm, cmd := m.breakingDescComponent.Update(msg, m)
		m.commit.breakingDesc = value
		m.footerComponent.textarea.SetValue("BREAKING CHANGE: " + value)
		return rm, cmd
	case scopeS:
		value, rm, cmd := m.scopeComponent.Update(msg, m)
		m.commit.scope = value
		return rm, cmd
	case descriptionS:
		value, rm, cmd := m.descriptionComponent.Update(msg, m)
		m.commit.description = value
		return rm, cmd
	case bodyS:
		value, rm, cmd := m.bodyComponent.Update(msg, m)
		m.commit.body = value
		return rm, cmd
	case footerS:
		value, rm, cmd := m.footerComponent.Update(msg, m)
		m.commit.footer = value
		return rm, cmd
	case previewS:
		value, rm, cmd := m.previewComponent.Update(msg, m)
		if !value {
			return rm, tea.Quit
		}
		return rm, cmd
	case commitS:
		cmd := m.commitMessage()
		return m, cmd
	}
	return m, nil
}

// View
func (m Model) View() string {
	switch m.state {
	case typeS:
		return style.Margin.Render(m.typeComponent.View())
	case breakingS:
		return style.Margin.Render(m.breakingComponent.View())
	case breakingDescS:
		return style.Margin.Render(m.breakingDescComponent.View())
	case scopeS:
		return style.Margin.Render(m.scopeComponent.View())
	case descriptionS:
		return style.Margin.Render(m.descriptionComponent.View())
	case bodyS:
		return style.Margin.Render(m.bodyComponent.View())
	case footerS:
		return style.Margin.Render(m.footerComponent.View())
	case previewS:
		return style.Margin.Render(m.previewComponent.View() + "\n\n" + m.previewCommit())
	case commitS:
		return "Commit View"
	default:
		return "not component view"
	}
}

// Utils
func (m Model) previewCommit() string {

	var commit string

	if !m.commit.breaking {
		commit = fmt.Sprintf(
			"%s%s: %s \n\n%s \n\n%s",
			m.commit.typeCommit,
			"("+m.commit.scope+")",
			m.commit.description,
			m.commit.body,
			m.commit.footer)
		return commit
	}

	commit = fmt.Sprintf(
		"%s%s!: %s \n\n%s\n\n%s",
		m.commit.typeCommit,
		"("+m.commit.scope+")",
		m.commit.description,
		m.commit.body,
		m.commit.footer,
	)

	return commit
}

func (m Model) commitMessage() tea.Cmd {

	fmt.Println(m.commit.scope)

	if m.commit.breaking {
		m.commit.scope = "(" + m.commit.scope + ")!: "
	} else {
		m.commit.scope = "(" + m.commit.scope + "): "
	}

	fmt.Printf(
		"git commit -m %s%s%s -m %s -m %s",
		m.commit.typeCommit,
		m.commit.scope,
		m.commit.description,
	)

	cmd := exec.Command(
		"git",
		"commit",
		"-m "+m.commit.typeCommit+m.commit.scope+m.commit.description,
		"-m "+m.commit.body, "-m "+m.commit.footer,
	)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: ", err)
	}
	fmt.Printf("output: %v", out)
	return nil
}
