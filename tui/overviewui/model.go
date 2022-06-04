package overviewui

import (
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/rwxd/zaster/internal"
	"github.com/rwxd/zaster/tui/constants"
	"github.com/rwxd/zaster/tui/models"
)

const (
	columnKeyID          = "id"
	columnKeyValue       = "value"
	columnKeyTime        = "time"
	columnKeyPayee       = "payee"
	columnKeyAccount     = "account"
	columnKeyCategory    = "category"
	columnKeyBudget      = "budget"
	columnKeyDescription = "description"
)

type SelectMsg struct {
	transaction models.TransactionModel
}

type Model struct {
	table       table.Model
	TableWidth  int
	TableMargin int
	selected    int
	cursor      int
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
		m.TableWidth = msg.Width
		m.recalculateTable()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Select):
			highlightedRow := m.table.HighlightedRow()
			selected := highlightedRow.Data[columnKeyID]
			log.Println("Selected transaction: ", selected)
		default:
			m.table, cmd = m.table.Update(msg)
		}
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}
	body.WriteString(m.table.View() + "\n")
	return body.String()
}

func (m *Model) recalculateTable() {
	m.table = m.table.WithTargetWidth(m.TableWidth - m.TableMargin)
}

func getFooter() string {
	var footer string
	keymaps := [][]string{
		{constants.Keymap.Up.Help().Key, constants.Keymap.Up.Help().Desc},
		{constants.Keymap.Down.Help().Key, constants.Keymap.Down.Help().Desc},
		{constants.Keymap.Create.Help().Key, constants.Keymap.Create.Help().Desc},
		{constants.Keymap.Edit.Help().Key, constants.Keymap.Edit.Help().Desc},
		{constants.Keymap.Delete.Help().Key, constants.Keymap.Delete.Help().Desc},
		{constants.Keymap.Quit.Help().Key, constants.Keymap.Quit.Help().Desc},
	}
	for index, keys := range keymaps {
		footer += keys[0] + " " + keys[1]
		if index+1 != len(keymaps) {
			footer += " · "
		}
	}
	// footer := fmt.Sprintf("%s up - %s down - %s create - %s edit - %s delete - %s quit",
	// 	constants.Keymap.Up.Help().Key,
	// 	constants.Keymap.Down.Help().Key,
	// 	constants.Keymap.Create.Help().Key,
	// 	constants.Keymap.Edit.Help().Key,
	// 	constants.Keymap.Delete.Help().Key,
	// 	constants.Keymap.Quit.Help().Key,
	// )
	return constants.HelpStyle.Render(footer)
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

func transformTransactionModelsToTableRows(transactionModels []models.TransactionModel) []table.Row {
	rows := make([]table.Row, len(transactionModels))
	for i, t := range transactionModels {
		rows[i] = table.NewRow(table.RowData{
			columnKeyValue:       t.Transaction.Value,
			columnKeyTime:        t.Transaction.Time.Format("02-01-2006"),
			columnKeyPayee:       t.Transaction.Payee,
			columnKeyAccount:     t.Transaction.Account,
			columnKeyCategory:    t.Transaction.Category,
			columnKeyBudget:      t.Transaction.Budget,
			columnKeyDescription: t.Transaction.Description,
			columnKeyID:          t.Transaction.ID.String(),
		})
	}
	return rows
}

func createTableColumns() []table.Column {
	return []table.Column{
		table.NewFlexColumn(columnKeyValue, "Value", 5),
		table.NewColumn(columnKeyTime, "Time", 10),
		table.NewFlexColumn(columnKeyPayee, "Payee", 15),
		table.NewFlexColumn(columnKeyAccount, "Account", 15),
		table.NewFlexColumn(columnKeyCategory, "Category", 15),
		table.NewFlexColumn(columnKeyBudget, "Budget", 15),
		table.NewFlexColumn(columnKeyDescription, "Description", 15),
	}
}

func NewOverviewModel() Model {
	tableRows := transformTransactionModelsToTableRows(loadTransactions())
	tableColumns := createTableColumns()

	keys := table.DefaultKeyMap()
	keys.RowUp.SetKeys(constants.Keymap.Up.Keys()...)
	keys.RowDown.SetKeys(constants.Keymap.Down.Keys()...)

	transactionTable := table.New(tableColumns).
		WithRows(tableRows).
		WithTargetWidth(60).
		WithKeyMap(keys).
		HeaderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#e07a5f")).Bold(true)).
		HighlightStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#81b29a")).Bold(true)).
		Focused(true).
		WithBaseStyle(
			lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("#f4f1de")).
				Foreground(lipgloss.Color("#e6e6fa")).
				Align(lipgloss.Left),
		).
		SortByAsc(columnKeyTime).
		WithStaticFooter(getFooter())

	m := Model{
		table: transactionTable,
	}
	return m
}
