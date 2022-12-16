package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Style Variable
var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4).PaddingTop(1)
	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("63")).
				PaddingTop(1)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {

	item, ok := listItem.(item)

	if !ok {
		return
	}

	str := fmt.Sprintf("%s", item)

	fn := itemStyle.Render

	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

// The struct of the list for bubble
type typeCommitModel struct {
	list     list.Model
	choice   string
	quitting bool
}

// Initial / Create the model

func newTypeModel() *typeCommitModel {
	items := []list.Item{
		item("🌟 feat"),
		item("🐛 fix"),
		item("🎉 init"),
		item("📝 docs"),
		item("🔨 refactor"),
		item("🏗️  build"),
		item("🎨 style"),
		item("✅ test"),
		item("🐎 perf"),
	}
	l := list.New(items, itemDelegate{}, 50, 15)
	l.Title = "Type of the commit"
	l.SetShowHelp(false)

	return &typeCommitModel{
		list: l,
	}
}


// The Update function
func (m *typeCommitModel) Update(msg tea.Msg, tm *Model) (string, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			item, ok := m.list.SelectedItem().(item)
			if ok {
				choice := string(item[2:])
				m.choice = strings.ReplaceAll(choice, " ", "") // Remove space
				tm.state++
			}
			return m.choice[2:], nil
		}

	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return "", cmd
}

func (m typeCommitModel) View() string {
	return "\n" + m.list.View()
}
