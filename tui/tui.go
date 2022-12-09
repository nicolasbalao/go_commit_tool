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
	bodyS
	scopeS
	descriptionS
	footerS
)

// Global struct of the app

type Model struct {
	typeComponent     *typeModel
	breakingComponent *breakingModel
	bodyComponent     *bodyModel

	/*Components
	  scopeComponent *scopeModel
	  descriptionComponent *descriptionModel
	  bodyComponent *bodyModel
	  footerComponent *footerModel
	*/

	commit commitMessage
	state  State
}

// Create the Model
func NewModel() Model {
	return Model{
		typeComponent:     newTypeModel(),
		breakingComponent: newBreakingModel(),
		bodyComponent:     newBodyModel(),
		state:             bodyS,
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
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	// Switch on the state of the app
	switch m.state {
	//Call update function of the components
	case typeS:
		return typeUpdate(msg, m)
	case breakingS:
		return breakingUpdate(msg, m)
	case bodyS:
		return m.bodyComponent.Update(msg, m)
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
	case bodyS:
		return m.bodyComponent.View()
	default:
		return ""
	}
}
