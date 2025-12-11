// Package connectui: TUI for wctl connect
package connectui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jwil007/wifictl/internal/connect"
)

type Model struct {
	Mode     AppMode
	Table    table.Model
	Form     FormModel
	Selected connect.SSIDEntry
}

type AppMode int

const (
	TableMode AppMode = iota
	FormMode
)

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
}

func (m model) View() string {
}

func Tui() {
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
