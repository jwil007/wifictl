package connectui

import tea "github.com/charmbracelet/bubbletea"

type formModel interface {
	Update(msg tea.Msg) (formModel, tea.Cmd)
	View() string
}

type openForm struct{}

func (f openForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f openForm) View() string {
	s := "PlaceholderOpen"
	return s
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

func newOWEForm() formModel {
	return oweForm{}
}

type pskForm struct{}

func (f pskForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f pskForm) View() string {
	s := "PlaceholderPSK"
	return s
}

func newPSKForm() formModel {
	return pskForm{}
}

type saeForm struct{}

func (f saeForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f saeForm) View() string {
	s := "PlaceholdersAE"
	return s
}

func newSAEForm() formModel {
	return saeForm{}
}

type eapForm struct{}

func (f eapForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f eapForm) View() string {
	s := "PlaceholderEAP"
	return s
}

func newEAPForm() formModel {
	return eapForm{}
}

type peapForm struct{}

func (f peapForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f peapForm) View() string {
	s := "Placeholder"
	return s
}

func newPEAPForm() formModel {
	return peapForm{}
}

type eapTLSForm struct{}

func (f eapTLSForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f eapTLSForm) View() string {
	s := "Placeholder"
	return s
}

func newEapTLSForm() formModel {
	return eapTLSForm{}
}
