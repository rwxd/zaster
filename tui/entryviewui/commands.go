package entryviewui

import tea "github.com/charmbracelet/bubbletea"

func backCmd() tea.Cmd {
	return func() tea.Msg {
		return BackMsg{}
	}
}
