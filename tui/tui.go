package tui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwxd/zaster/internal"
	"github.com/rwxd/zaster/tui/entryviewui"
	"github.com/rwxd/zaster/tui/overviewui"
)

type sessionState int

const overviewViewState sessionState = 1
const transactionViewState sessionState = 2

type MainModel struct {
	overviewModel     tea.Model
	transactionModel  tea.Model
	state             sessionState
	db                *internal.JSONDatabase
	activeTransaction internal.Transaction
}

// Init run any intial IO on program start
func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case overviewui.SelectMsg:
		m.state = transactionViewState
		m.activeTransaction = msg.ActiveTransaction
	case entryviewui.BackMsg:
		m.state = overviewViewState
	}

	switch m.state {
	case overviewViewState:
		newOverview, newCmd := m.overviewModel.Update(msg)
		overviewModel := newOverview.(overviewui.Model)
		m.overviewModel = overviewModel
		cmd = newCmd
	case transactionViewState:
		newTransactionModel, newCmd := m.transactionModel.Update(msg)
		transactionModel := newTransactionModel.(entryviewui.Model)
		m.transactionModel = transactionModel
		cmd = newCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case transactionViewState:
		m.transactionModel = entryviewui.NewEntryViewModel(m.activeTransaction, m.db)
		return m.transactionModel.View()
	default:
		return m.overviewModel.View()
	}
}

func NewMainModel(db *internal.JSONDatabase) MainModel {
	return MainModel{
		state:            overviewViewState,
		overviewModel:    overviewui.NewOverviewModel(db),
		transactionModel: entryviewui.NewEntryViewModel(internal.Transaction{}, db),
	}
}

// StartTea the entry point for the UI. Initializes the model.
func StartTea(db *internal.JSONDatabase) {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	m := NewMainModel(db)
	m.state = overviewViewState
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running programm:", err)
		os.Exit(1)
	}
}
