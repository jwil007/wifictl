package connectui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type formModel interface {
	Update(msg tea.Msg) (formModel, tea.Cmd)
	View() string
}

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
			Padding(1).
			Height(2).
			Width(70),

		Footer: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Bold(true),
	}
}

func sendMsg(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

type openForm struct{}

func (f openForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f openForm) View() string {
	s := "PlaceholderOpen"
	return s
}

func (f openForm) Submit() interface{} {
	return nil
}

func newOpenForm() formModel {
	return openForm{}
}

type oweForm struct{}

func (f oweForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f oweForm) View() string {
	s := "PlaceholderOWE"
	return s
}

func (f oweForm) Submit() interface{} {
	return nil
}

func newOWEForm() formModel {
	return oweForm{}
}

type pskForm struct {
	TextInput textinput.Model
	Err       error
	Styles    formStyles
}

type pskSubmitMsg struct {
	Passphrase string
}

func (f pskForm) Init() tea.Cmd {
	return textinput.Blink
}

func (f pskForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			msg := f.Submit()
			return f, sendMsg(msg)
		}
		f.TextInput, cmd = f.TextInput.Update(msg)
		return f, cmd
	}
	return f, nil
}

func (f pskForm) View() string {
	title := f.Styles.Title.Render("Connect to SSID")
	textBox := f.Styles.TextBox.Render(f.TextInput.View())
	footer := f.Styles.Footer.Render("enter to submit | esc to close")

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		textBox,
		footer,
	)
	return f.Styles.Container.Render(content)
}

func (f pskForm) Submit() pskSubmitMsg {
	return pskSubmitMsg{
		Passphrase: f.TextInput.Value(),
	}
}

func newPSKForm() formModel {
	ti := textinput.New()
	ti.Placeholder = "Enter passphrase..."
	ti.PlaceholderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
	ti.Focus()
	ti.CharLimit = 64
	return pskForm{
		TextInput: ti,
		Err:       nil,
		Styles:    defaultFormStyles(),
	}
}

type saeForm struct {
	textInput textinput.Model
	err       error
}

type saeSubmitMsg struct {
	Passphrase string
}

func (f saeForm) Init() tea.Cmd {
	return textinput.Blink
}

func (f saeForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			msg := f.Submit()
			return f, sendMsg(msg)
		}
		f.textInput, cmd = f.textInput.Update(msg)
		return f, cmd
	}
	return f, nil
}

func (f saeForm) View() string {
	return fmt.Sprintf("Enter passphrase\n\n%s\n\n%s",
		f.textInput.View(),
		"enter to submit\nesc to close",
	)
}

func (f saeForm) Submit() saeSubmitMsg {
	return saeSubmitMsg{
		Passphrase: f.textInput.Value(),
	}
}

func newSAEForm() formModel {
	ti := textinput.New()
	ti.Placeholder = "Passphrase..."
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 70
	return saeForm{
		textInput: ti,
		err:       nil,
	}
}

type eapForm struct{}

func (f eapForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f eapForm) View() string {
	s := "PlaceholderEAP"
	return s
}

func (f eapForm) Submit() interface{} {
	return nil
}

func newEAPForm() formModel {
	return eapForm{}
}

//type peapForm struct{}
//
//func (f peapForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
//	return nil, nil
//}
//
//func (f peapForm) View() string {
//	s := "Placeholder"
//	return s
//}
//
//func newPEAPForm() formModel {
//	return peapForm{}
//}
//
//type eapTLSForm struct{}
//
//func (f eapTLSForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
//	return nil, nil
//}
//
//func (f eapTLSForm) View() string {
//	s := "Placeholder"
//	return s
//}
//
//func newEapTLSForm() formModel {
//	return eapTLSForm{}
//}
