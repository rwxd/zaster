package overviewui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func selectTransactionCmd(transactionId string) tea.Cmd {
	return func() tea.Msg {
		log.Println("Selected transaction: ", transactionId)
		return SelectMsg{ActiveTransaction: transactionId}
	}
}
