package connectui

import tea "github.com/charmbracelet/bubbletea"

type FormModel interface {
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type openForm struct{}

func (f openForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (f openForm) View() string {
	s := "Placeholder"
	return s
}

type passphraseForm struct{}

func (f passphraseForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (f passphraseForm) View() string {
	s := "Placeholder"
	return s
}

type peapForm struct{}

func (f peapForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (f peapForm) View() string {
	s := "Placeholder"
	return s
}

type eapTLSForm struct{}

func (f eapTLSForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (f eapTLSForm) View() string {
	s := "Placeholder"
	return s
}
