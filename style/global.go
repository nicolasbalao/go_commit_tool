package style

import "github.com/charmbracelet/lipgloss"

var (
	Margin = lipgloss.NewStyle().Margin(1, 2).Width(70)

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

    HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("60")).MarginTop(1)

    BorderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("63"))

    SubtitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))

    ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)
