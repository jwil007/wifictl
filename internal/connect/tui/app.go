// Package connectui: TUI for wctl connect
package connectui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jwil007/wifictl/internal/connect"
)

type Model struct {
	Mode     AppMode
	Table    tableModel
	Form     FormModel
	Selected connect.SSIDEntry
}

type AppMode int

const (
	TableMode AppMode = iota
	FormMode
)

func (m Model) InitialModel(ssidList []connect.SSIDEntry) Model {
	m.Mode = TableMode
	newTableModel(ssidList)
	return m
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.Mode {
	case TableMode:
		return m.Table.Update(msg)
	case FormMode:
		return m.Form.Update(msg)
	default:
		return nil, nil
	}
}

func (m Model) View() string {
	switch m.Mode {
	case TableMode:
		return m.Table.View()
	case FormMode:
		return m.Form.View()
	default:
		return "You shouldn't see this"
	}
}

func Tui() {
	var m Model
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
