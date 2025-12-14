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
	SSIDList []connect.SSIDEntry
	Height   int
	Width    int
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
type connectErrorMsg struct {
	Err error
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

func detectSecType(index int, ssidList []connect.SSIDEntry) string {
	sec := strings.Join(ssidList[index].SecType, "")

	switch {
	case sec == "":
		return "open"
	case strings.Contains(sec, "OWE"):
		return "owe"
	case strings.Contains(sec, "PSK"):
		return "psk"
	case strings.Contains(sec, "SAE"):
		return "sae"
	case strings.Contains(sec, "EAP"):
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

func doConnectCmd(iface string,
	ssidEntry connect.SSIDEntry,
	sec connect.WiFiSecurity,
) tea.Cmd {
	return func() tea.Msg {
		err := connect.DoConnect(iface, ssidEntry, sec)
		if err != nil {
			return connectErrorMsg{Err: err}
		}
		return nil
	}
}

func (m Model) Init() tea.Cmd {
	return doScanCmd(m.Iface)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			if m.Mode == tableMode {
				m.Mode = loadingMode
				return m, doScanCmd(m.Iface)
			}
		case "esc":
			if m.Mode == formMode {
				m.Mode = tableMode
				return m, nil
			}
		}

	case scanResultsMsg:
		m.SSIDList = msg.SSIDList
		rows := makeRows(msg.SSIDList)
		m.Table.table.SetRows(rows)
		m.Mode = tableMode
		return m, nil

	case selectedRowMsg:
		m.Selected = m.SSIDList[msg.Cursor]
		switch detectSecType(msg.Cursor, m.SSIDList) {
		case "open":
			m.Form = newOpenForm()
		case "owe":
			m.Form = newOWEForm()
		case "psk":
			m.Form = newPSKForm(m.Selected.SSID, false)
		case "sae":
			m.Form = newPSKForm(m.Selected.SSID, true)
		case "eap":
			m.Form = newEAPForm()
		default:
			m.Form = newOpenForm()
		}
		m.Mode = formMode
		return m, nil

	case pskSubmitMsg:
		var sec connect.PSKSec
		sec.Passphrase = msg.Passphrase
		sec.SAE = msg.SAE
		cmd := doConnectCmd(m.Iface, m.Selected, sec)
		return m, cmd

	case connectErrorMsg:
		fmt.Println("connect error", msg.Err)
	}

	switch m.Mode {
	case tableMode:
		newTable, cmd := m.Table.Update(msg)
		m.Table = newTable
		return m, cmd
	case formMode:
		newForm, cmd := m.Form.Update(msg)
		m.Form = newForm
		return m, cmd
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
	case loadingMode:
		return "loading..."
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
