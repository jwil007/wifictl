package connectui

import tea "github.com/charmbracelet/bubbletea"

type FormModel interface {
	Update()
	View()
}

type openForm struct{}

func (f openForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
}

func (f openForm) View() string {
}

type passphraseForm struct{}

func (f passphraseForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
}

func (f passphraseForm) View() string {
}

type peapForm struct{}

func (f peapForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
}

func (f peapForm) View() string {
}

type eapTLSForm struct{}

func (f eapTLSForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
}

func (f eapTLSForm) View() string {
}
