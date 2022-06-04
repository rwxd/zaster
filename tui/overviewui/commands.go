package overviewui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwxd/zaster/tui/models"
)

func selectTransactionCmd(transaction models.TransactionModel) tea.Cmd {
	return func() tea.Msg {
		log.Println("Selected transaction: ", transaction.Title())
		return SelectMsg{transaction: transaction}
	}
}
