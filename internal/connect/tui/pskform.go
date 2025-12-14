package connectui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pskForm struct {
	TextInput textinput.Model
	Err       error
	Styles    formStyles
	SSID      string
	SAE       bool
}

type pskSubmitMsg struct {
	Passphrase string
	SAE        bool
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
	title := f.Styles.Title.Render("Connect to " + f.SSID)
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
		SAE:        f.SAE,
	}
}

func newPSKForm(ssid string, sae bool) formModel {
	ti := textinput.New()
	ti.Placeholder = "Enter passphrase..."
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 70

	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '*'

	return pskForm{
		TextInput: ti,
		Err:       nil,
		Styles:    defaultFormStyles(),
		SSID:      ssid,
		SAE:       sae,
	}
}
