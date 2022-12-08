package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nicolasbalao/go_commit_toll/tui"
)

func main() {

	p := tea.NewProgram(tui.NewModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
