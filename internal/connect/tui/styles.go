package connectui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

type formStyles struct {
	Container lipgloss.Style
	Title     lipgloss.Style
	TextBox   lipgloss.Style
	Footer    lipgloss.Style
	Error     lipgloss.Style
}

func defaultFormStyles() formStyles {
	return formStyles{
		Container: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1, 2).
			Width(80),

		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("69")),

		TextBox: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1).
			Width(70),

		Footer: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Bold(true),
	}
}

func applyTextInputStyles(ti *textinput.Model) {
	ti.TextStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("15"))

	ti.PlaceholderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	ti.PromptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252"))

	ti.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("69"))
}
