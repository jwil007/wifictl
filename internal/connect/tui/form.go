package connectui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type formModel interface {
	Update(msg tea.Msg) (formModel, tea.Cmd)
	View() string
}

func sendMsg(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
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
