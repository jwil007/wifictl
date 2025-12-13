// Package connectui: TUI for wctl connect
package connectui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jwil007/wifictl/internal/connect"
)

type Model struct {
	Mode     appMode
	Table    tableModel
	Form     formModel
	Selected connect.SSIDEntry
	Iface    string
}

type appMode int

const (
	loadingMode appMode = iota
	tableMode
	formMode
)

type scanErrorMsg struct {
	Err error
}
type scanResultsMsg struct {
	SSIDList []connect.SSIDEntry
}

func initialModel() Model {
	return Model{
		Mode: loadingMode,
		// hard coding iface for n
		Iface: "wlp0s20f3",
		Form:  nil,
		Table: newTableModel(),
	}
}

func detectSecType(row []string) string {
	switch {
	case row[3] == "":
		return "open"
	case strings.Contains(row[3], "OWE"):
		return "owe"
	case strings.Contains(row[3], "PSK"):
		return "psk"
	case strings.Contains(row[3], "SAE"):
		return "sae"
	case strings.Contains(row[3], "EAP"):
		return "eap"
	default:
		return "invalid sec type"
	}
}

func doScanCmd(iface string) tea.Cmd {
	return func() tea.Msg {
		ssidList, err := connect.DoScan(iface)
		if err != nil {
			return scanErrorMsg{Err: err}
		}
		return scanResultsMsg{SSIDList: ssidList}
	}
}

func (m Model) Init() tea.Cmd {
	return doScanCmd(m.Iface)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case scanResultsMsg:
		rows := makeRows(msg.SSIDList)
		m.Table.table.SetRows(rows)
		m.Mode = tableMode
		return m, nil

	case selectedRowMsg:
		switch detectSecType(msg.Row) {
		case "open":
			m.Form = newOpenForm()
		case "owe":
			m.Form = newOWEForm()
		case "psk":
			m.Form = newPSKForm()
		case "sae":
			m.Form = newSAEForm()
		case "eap":
			m.Form = newEAPForm()
		}
		m.Mode = formMode
		return m, nil
	}

	switch m.Mode {
	case tableMode:
		newTable, cmd := m.Table.Update(msg)
		m.Table = newTable
		return m, cmd
	case formMode:
		return m, nil
	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.Mode {
	case tableMode:
		return m.Table.View()
	case formMode:
		return m.Form.View()
	default:
		return "You shouldn't see this"
	}
}

func Tui() {
	m := initialModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
