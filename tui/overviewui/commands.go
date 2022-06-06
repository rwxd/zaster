package overviewui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwxd/zaster/internal"
)

func selectTransactionCmd(transaction internal.Transaction) tea.Cmd {
	return func() tea.Msg {
		log.Println("Selected transaction: ", transaction.Str())
		return SelectMsg{ActiveTransaction: transaction}
	}
}
