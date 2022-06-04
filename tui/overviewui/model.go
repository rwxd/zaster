package overviewui

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rwxd/zaster/internal"
	"github.com/rwxd/zaster/tui/constants"
	"github.com/rwxd/zaster/tui/models"
)

type SelectMsg struct {
	transaction models.TransactionModel
}

type Model struct {
	list     list.Model
	selected int
	cursor   int
}

// Load transactions database on start
func (m Model) Init() tea.Cmd {
	// TODO: Load database
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Select):
			cmd = selectTransactionCmd(m.getActiveTransaction())
			m.selected = m.cursor
		default:
			m.list, cmd = m.list.Update(msg)
		}
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return constants.DocStyle.Render(m.list.View() + "\n")
}

func (m Model) getActiveTransaction() models.TransactionModel {
	items := m.list.Items()
	activeItem := items[m.list.Index()]
	return activeItem.(models.TransactionModel)
}

func transactionModelsToListItems(transactions []models.TransactionModel) []list.Item {
	items := make([]list.Item, len(transactions))
	for i, transaction := range transactions {
		items[i] = list.Item(transaction)
	}
	return items
}

func loadTransactions() []models.TransactionModel {
	transaction1, _ := internal.NewTransaction(25.0, time.Now(), "Marten", "Commerzbank", "", "", "Essen", internal.MoneyInflow)
	transaction2, _ := internal.NewTransaction(36.99, time.Now(), "Peter", "Commerzbank", "", "", "", internal.MoneyOutflow)
	transaction3, _ := internal.NewTransaction(29.72, time.Now(), "Versicherung", "Commerzbank", "", "", "Rückzahlung", internal.MoneyInflow)
	transaction4, _ := internal.NewTransaction(129.53, time.Now(), "DB", "Commerzbank", "", "", "9€ Ticket", internal.MoneyOutflow)
	transaction5, _ := internal.NewTransaction(12.0, time.Now(), "Mc Donalds", "Commerzbank", "", "", "Essen gehen", internal.MoneyOutflow)
	transaction6, _ := internal.NewTransaction(2502.0, time.Now(), "Firma", "", "", "", "Gehalt", internal.MoneyInflow)
	tempTransactions := []internal.Transaction{transaction1, transaction2, transaction3, transaction4, transaction5, transaction6}
	transactionModels := make([]models.TransactionModel, len(tempTransactions))
	for i, transaction := range tempTransactions {
		transactionModels[i] = models.NewTransactionModel(transaction)
	}
	return transactionModels
}

func NewOverviewModel() Model {
	items := transactionModelsToListItems(loadTransactions())

	m := Model{
		list: list.NewModel(items, list.NewDefaultDelegate(), 0, 0),
	}

	m.list.Title = "Transactions"
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			constants.Keymap.Select,
			constants.Keymap.Quit,
			constants.Keymap.Create,
			constants.Keymap.Delete,
		}
	}
	return m
}
