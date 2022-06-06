package entryviewui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwxd/zaster/internal"
	"github.com/rwxd/zaster/tui/constants"
)

type BackMsg struct {
}

type Model struct {
	db                *internal.JSONDatabase
	ActiveTransaction internal.Transaction
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "q":
			cmd = backCmd()
		case key.Matches(msg, constants.Keymap.Back):
			cmd = backCmd()
		case key.Matches(msg, constants.Keymap.Quit):
			return m, tea.Quit
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}
	body.WriteString("Viewing at transcation: ")
	body.WriteString(m.ActiveTransaction.Str() + "\n")
	return body.String()
}

func NewEntryViewModel(transaction internal.Transaction, db *internal.JSONDatabase) Model {
	return Model{
		ActiveTransaction: transaction,
		db:                db,
	}
}
