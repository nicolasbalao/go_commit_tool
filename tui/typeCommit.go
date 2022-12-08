package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
)

// Style Variable
var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
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
type typeModel struct {
	list     list.Model
	choice   string
	quitting bool
}

// Initial / Create the model

func newTypeModel() *typeModel {
	items := []list.Item{
		item("feat"),
		item("fix"),
		item("docs"),
		item("refactor"),
		item("build"),
		item("style"),
		item("test"),
		item("chore"),
		item("perf"),
	}

	const defaultWidth = 40

	l := list.New(items, itemDelegate{}, defaultWidth, 15)
	l.Title = "Type of the commit"

	return &typeModel{
		list: l,
	}
}

// Init function, rly need?
func (m typeModel) Init() tea.Cmd {
	return nil
}

// The Update function

func typeUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
        case "enter":
            item, ok := m.typeComponent.list.SelectedItem().(item)
            if ok {
                m.typeComponent.choice = string(item)
                m.commit.scope = string(item)
                m.state++
            }
            return m, nil
		}

	}
    var cmd tea.Cmd
    m.typeComponent.list, cmd = m.typeComponent.list.Update(msg)

    return m, cmd
}

func (m typeModel) View() string{
    return "\n" + m.list.View()
}
