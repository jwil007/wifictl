package connectui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type openForm struct {
	Styles formStyles
	SSID   string
	OWE    bool
}

type openSubmitMsg struct {
	OWE bool
}

func (f openForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			msg := f.Submit()
			return f, sendMsg(msg)
		}
		return f, cmd
	}
	return f, nil
}

func (f openForm) View() string {
	title := f.Styles.Title.Render("Connect to " + f.SSID)
	footer := f.Styles.Footer.Render("enter to connect | esc to close")
	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		"Open SSID - no config needed",
		"",
		footer,
	)
	return f.Styles.Container.Render(content)
}

func (f openForm) Submit() openSubmitMsg {
	return openSubmitMsg{
		OWE: f.OWE,
	}
}

func newOpenForm(ssid string, owe bool) formModel {
	return openForm{
		Styles: defaultFormStyles(),
		SSID:   ssid,
		OWE:    owe,
	}
}
