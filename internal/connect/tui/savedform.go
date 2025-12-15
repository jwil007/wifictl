package connectui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type savedForm struct {
	Styles formStyles
	SSID   string
}

type savedSubmitMsg struct {
	Connect bool
	Forget  bool
}

func (f savedForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			msg := f.Submit("connect")
			return f, sendMsg(msg)
		case "F":
			msg := f.Submit("forget")
			return f, sendMsg(msg)
		}
		return f, cmd
	}
	return f, nil
}

func (f savedForm) View() string {
	title := f.Styles.Title.Render("Connect to " + f.SSID)
	footer := f.Styles.Footer.Render("enter to connect | F (shift+f) to forget SSID | esc to close")

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		"SSID configuration saved",
		"",
		footer,
	)
	return f.Styles.Container.Render(content)
}

func (f savedForm) Submit(option string) savedSubmitMsg {
	if option == "connect" {
		return savedSubmitMsg{
			Connect: true,
			Forget:  false,
		}
	}
	if option == "forget" {
		return savedSubmitMsg{
			Connect: false,
			Forget:  true,
		}
	}
	return savedSubmitMsg{}
}

func newSavedForm(ssid string) formModel {
	return savedForm{
		Styles: defaultFormStyles(),
		SSID:   ssid,
	}
}
