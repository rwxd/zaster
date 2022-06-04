package entryui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	transaction string
}

func (m Model) Init() tea.Cmd {
	log.Println("Initializing entry model")
	return nil
}

func NewEntryModel(transaction string) Model {
	log.Printf("New entry model with transaction %s\n", transaction)
	return Model{
		transaction: transaction,
	}
}
