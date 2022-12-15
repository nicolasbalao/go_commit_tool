package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 70
)

type tickMsg time.Time

type progressModel struct {
	progress progress.Model
}

func newProgressModel() *progressModel {
	return &progressModel{
		progress: progress.New(progress.WithDefaultGradient()),
	}
}

func (m *progressModel) Init() tea.Cmd {
	return tickCmd()
}

func (m *progressModel) Update(msg tea.Msg, tm *Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return nil
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return cmd
	default:
		if m.progress.Percent() == 1.00 {
            tm.state++
			return nil
		}
		cmd := m.progress.IncrPercent(0.25)
		return tea.Batch(tickCmd(), cmd)
	}

}

func (m progressModel) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.View() + "\n\n" +
		pad
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
