package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Structures and enums

type commitMessage struct {
	typeCommit  string
	scope       string // optional
	description string
	body        string // optional
	breaking    bool
	footer      string // optional
}

// Enums for the state of the app
type State int

const (
	typeS State = iota
	breakingS
	scopeS
	descriptionS
	bodyS
	footerS
	commitS
)

// Global struct of the app

type Model struct {
	typeComponent        *typeModel
	breakingComponent    *breakingModel
	scopeComponent       *textInputModel
	descriptionComponent *textInputModel
	bodyComponent        *textAreaModel
	footerComponent      *textAreaModel

    focusedTextArea bool
	commit *commitMessage
	state  State
}

// Create the Model
func NewModel() Model {

	commitMessae := commitMessage{
		scope: "default",
	}

	return Model{
		typeComponent:        newTypeModel(),
		breakingComponent:    newBreakingModel(),
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
		commit: &commitMessae,
		state:  scopeS,
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
            if !m.focusedTextArea{
                return m, tea.Quit
            }
		}
         
	}
	// Switch on the state of the app
	switch m.state {
	//Call update function of the components
	case typeS:
		return typeUpdate(msg, m)
	case breakingS:
		return breakingUpdate(msg, m)
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

	case commitS:
		m.printCommit()
		return m, nil
	}
	return m, nil
}

// View
func (m Model) View() string {
	switch m.state {
	case typeS:
		return m.typeComponent.View()
	case breakingS:
		return m.breakingComponent.View()
	case scopeS:
		return m.scopeComponent.View()
	case descriptionS:
		return m.descriptionComponent.View()
	case bodyS:
		return m.bodyComponent.View()
	case footerS:
		return m.footerComponent.View()
	case commitS:
		m.printCommit()
		return "Commit View"
	default:
		return "not component view"
	}
}

// Utils
func (m Model) printCommit() {
	fmt.Printf(
		"breaking: %v \ntype: %s \nscope: %s\ndescription: %s\nbody: %s\nfooter: %s\n",
		m.commit.breaking,
		m.commit.typeCommit,
		m.commit.scope,
		m.commit.description,
		m.commit.body,
		m.commit.footer,
	)
}
