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
