package connectui

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jwil007/wifictl/internal/connect"
)

type tableModel struct {
	table table.Model
}

type selectedRowMsg struct {
	Row []string
}

func makeRows(ssidList []connect.SSIDEntry) []table.Row {
	var rows []table.Row
	for _, entry := range ssidList {
		row := table.Row{
			entry.SSID,
			strconv.Itoa(entry.RSSI),
			strconv.Itoa(entry.BSSIDCount),
			strings.Join(entry.SecType, " "),
			strings.Join(entry.Bands, " "),
		}
		rows = append(rows, row)
	}
	return rows
}

func sendSelected(row []string) tea.Cmd {
	return func() tea.Msg {
		return selectedRowMsg{Row: row}
	}
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, sendSelected(m.table.SelectedRow())

		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func newTableModel(ssidList []connect.SSIDEntry) tableModel {
	columns := []table.Column{
		{Title: "SSID", Width: 14},
		{Title: "RSSI", Width: 4},
		{Title: "AP Ct.", Width: 8},
		{Title: "Security", Width: 20},
		{Title: "Bands", Width: 20},
	}
	rows := makeRows(ssidList)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(14),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return tableModel{table: t}
}
