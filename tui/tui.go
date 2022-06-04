package tui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwxd/zaster/tui/overviewui"
)

type sessionState int

const overviewViewState sessionState = 1
const entryViewState sessionState = 2

type MainModel struct {
	overview tea.Model
	entry    tea.Model
	state    sessionState
}

// Init run any intial IO on program start
func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.state = overviewViewState

	switch m.state {
	case overviewViewState:
		newTransaction, newCmd := m.overview.Update(msg)
		overviewModel := newTransaction.(overviewui.Model)
		m.overview = overviewModel
		cmd = newCmd
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

func NewMainModel() MainModel {
	return MainModel{
		state:    overviewViewState,
		overview: overviewui.NewOverviewModel(),
	}
}

// StartTea the entry point for the UI. Initializes the model.
func StartTea() {
	if os.Getenv("DEBUG") != "" {
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
	}

	m := NewMainModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running programm:", err)
		os.Exit(1)
	}
}
