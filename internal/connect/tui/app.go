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
	Iface    string
}

type AppMode int

const (
	LoadingMode AppMode = iota
	TableMode
	FormMode
)

type ScanErrorMsg struct {
	Err error
}
type ScanResultsMsg struct {
	SSIDList []connect.SSIDEntry
}

func InitialModel() Model {
	return Model{
		Mode: LoadingMode,
		// hard coding iface for n
		Iface: "wlp0s20f3",
		Form:  nil,
		Table: NewTableModel(),
	}
}

func DoScanCmd(iface string) tea.Cmd {
	return func() tea.Msg {
		ssidList, err := connect.DoScan(iface)
		if err != nil {
			return ScanErrorMsg{Err: err}
		}
		return ScanResultsMsg{SSIDList: ssidList}
	}
}

func (m Model) Init() tea.Cmd {
	return DoScanCmd(m.Iface)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ScanResultsMsg:
		rows := makeRows(msg.SSIDList)
		m.Table.table.SetRows(rows)
		m.Mode = TableMode
		return m, nil
	case selectedRowMsg:
		m.Mode = FormMode
		return m, nil

	}

	switch m.Mode {
	case TableMode:
		return m.Table.Update(msg)
	case FormMode:
		return m.Form.Update(msg)
	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.Mode {
	case TableMode:
		return m.Table.View()
	case FormMode:
		return "you are in form mode"
	default:
		return "You shouldn't see this"
	}
}

func Tui() {
	m := InitialModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
