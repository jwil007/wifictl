package connectui

import tea "github.com/charmbracelet/bubbletea"

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
