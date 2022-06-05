package tui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwxd/zaster/internal"
	"github.com/rwxd/zaster/tui/overviewui"
)

type sessionState int

const overviewViewState sessionState = 1
const transactionViewState sessionState = 2

type MainModel struct {
	overview          tea.Model
	entry             tea.Model
	state             sessionState
	db                internal.JSONDatabase
	activeTransaction string
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
	}

	switch m.state {
	case overviewViewState:
		newOverview, newCmd := m.overview.Update(msg)
		overviewModel := newOverview.(overviewui.Model)
		m.overview = overviewModel
		cmd = newCmd
	case transactionViewState:
		// newTransaction :=
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	// switch m.state {
	// case entryViewState:
	// 	return m.entry.View()
	// default:
	// 	return m.overview.View()
	// }
	return m.overview.View()
}

func NewMainModel(db *internal.JSONDatabase) MainModel {
	return MainModel{
		state:    overviewViewState,
		overview: overviewui.NewOverviewModel(db),
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
