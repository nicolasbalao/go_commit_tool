/*
author: nicolas balao

This component handle update and view of other component with state.
*/
package tui

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nicolasbalao/go_commit_tool/style"
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
	sendCommitS
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
	progressComponent     *progressModel

	focusedTextArea bool
	commit          *commitMessage
	state           State
	helper          string
}

// Create the Model
func NewModel() Model {

	commitMessage := commitMessage{}

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
		progressComponent: newProgressModel(),

		previewComponent: newConfirmComponent("Preview", "Commit ?"),
		commit:           &commitMessage,
		state:            typeS,
        helper:           "ctrl+h/p: back  ctrl+l/n: next  enter: confirm",
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
		case "ctrl+n", "ctrl+l":
			if m.state != 7 {
				m.state++
			}
			return m, nil
		case "ctrl+p", "ctrl+h":
			m.state--
			if m.state < 0 {
				m.state = 0
			}
			return m, nil
		}

	}
	// Switch on the state of the app
	switch m.state {
	//Call update function of the components
	case typeS:
		value, cmd := m.typeComponent.Update(msg, &m)
		m.commit.typeCommit = value
		return m, cmd
	case breakingS:
		value, cmd := m.breakingComponent.Update(msg, &m)
		m.commit.breaking = value
		return m, cmd
	case breakingDescS:
		value, cmd := m.breakingDescComponent.Update(msg, &m)
		m.commit.breakingDesc = value
		m.footerComponent.textarea.SetValue("BREAKING CHANGE: " + value)
		return m, cmd
	case scopeS:
		value, cmd := m.scopeComponent.Update(msg, &m)
		m.commit.scope = value
		return m, cmd
	case descriptionS:
		value, cmd := m.descriptionComponent.Update(msg, &m)
		m.commit.description = value
		return m, cmd
	case bodyS:
		value, cmd := m.bodyComponent.Update(msg, &m)
		m.commit.body = value
		return m, cmd
	case footerS:
		value, cmd := m.footerComponent.Update(msg, &m)
		m.commit.footer = value
		return m, cmd
	case previewS:
        m.focusedTextArea = false
		if m.commit.typeCommit == "" || m.commit.description == "" {
			return m, nil
		}
		_, cmd := m.previewComponent.Update(msg, &m)
		return m, cmd
	case commitS:
        cmd := m.progressComponent.Update(msg, &m)
        return m, cmd
	case sendCommitS:
        cmd := m.sendCommitMesage()
		return m, cmd
	default:
		return m, tea.Quit
	}
}

// View
func (m Model) View() string {
	switch m.state {
	case typeS:
		return style.Margin.Render(m.typeComponent.View() + "\n" + style.HelpStyle.Render(m.helper))
	case breakingS:
		return style.Margin.Render(
			m.breakingComponent.View() + "\n" + style.HelpStyle.Render(m.helper),
		)
	case breakingDescS:
		return style.Margin.Render(
			m.breakingDescComponent.View() + "\n" + style.HelpStyle.Render(m.helper),
		)
	case scopeS:
		return style.Margin.Render(
			m.scopeComponent.View() + "\n" + style.HelpStyle.Render(m.helper),
		)
	case descriptionS:
		return style.Margin.Render(
			m.descriptionComponent.View() + "\n" + style.HelpStyle.Render(m.helper),
		)
	case bodyS:
		return style.Margin.Render(m.bodyComponent.View() + "\n" + style.HelpStyle.Render(m.helper))
	case footerS:
		return style.Margin.Render(
			m.footerComponent.View() + "\n" + style.HelpStyle.Render(m.helper),
		)
	case previewS:

		if m.commit.typeCommit == "" || m.commit.description == "" {
			return style.ErrorStyle.Render("Missing value scope or description")
		}

		return style.Margin.Render(
			m.previewComponent.View() + "\n\n" + style.SubtitleStyle.Render(
				"Commit Message",
			) + "\n" + style.BorderStyle.Render(
				m.previewCommit(),
			) + "\n" + style.HelpStyle.Render(
				m.helper,
			),
		)
	case commitS:
		return "Press enter for start sending" + "\n\n" + m.progressComponent.View()
	default:
		return "Bye Bye"
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

func (m Model) sendCommitMesage() tea.Cmd {

	git_icon := map[string]string{
		"feat":     ":sparkles: ",
		"init":     ":tada: ",
		"fix":      ":bug: ",
		"docs":     ":books: ",
		"refactor": ":hammer: ",
		"build":    ":construction: ",
		"style":    ":art: ",
		"test":     " :white_check_mark: ",
		"perf":     ":racehorse: ",
	}

	icon_message := git_icon[m.commit.typeCommit]

	if m.commit.breaking {
		m.commit.scope = "(" + m.commit.scope + ")!: "
	} else if m.commit.scope != "" {
		m.commit.scope = "(" + m.commit.scope + "): "
	} else {
		m.commit.typeCommit += ": "
	}

	cmd := exec.Command(
		"git",
		"commit",
		"-m "+icon_message+m.commit.typeCommit+m.commit.scope+m.commit.description,
		"-m "+m.commit.body, "-m "+m.commit.footer,
	)

	err := cmd.Run()

	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	return tea.Quit
}
