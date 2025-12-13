package connectui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type formModel interface {
	Update(msg tea.Msg) (formModel, tea.Cmd)
	View() string
	Submit() interface{}
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
	textInput textinput.Model
	err       error
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
			f.Submit()
			return f, nil
		}
		f.textInput, cmd = f.textInput.Update(msg)
		return f, cmd
	}
	return f, nil
}

func (f pskForm) View() string {
	return fmt.Sprintf("Enter passphrase\n\n%s\n\n%s",
		f.textInput.View(),
		"enter to submit\nesc to close",
	)
}

func (f pskForm) Submit() interface{} {
	return f.textInput.Value()
}

func newPSKForm() formModel {
	ti := textinput.New()
	ti.Placeholder = "Passphrase..."
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 70
	return pskForm{
		textInput: ti,
		err:       nil,
	}
}

type saeForm struct{}

func (f saeForm) Update(msg tea.Msg) (formModel, tea.Cmd) {
	return nil, nil
}

func (f saeForm) View() string {
	s := "PlaceholdersAE"
	return s
}

func (f saeForm) Submit() interface{} {
	return nil
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
