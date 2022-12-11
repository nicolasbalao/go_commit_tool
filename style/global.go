package style

import "github.com/charmbracelet/lipgloss"

var (
	Margin = lipgloss.NewStyle().Margin(1, 2).Width(50)

	TitleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1).
			MarginBottom(1)

	InputTextStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("62")).
			UnsetAlignHorizontal()

	InputPrompStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62"))
)
