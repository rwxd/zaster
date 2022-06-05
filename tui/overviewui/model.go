package overviewui

import (
	"fmt"
	"log"
	"strings"

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
	columnKeyBudget      = "budget"
	columnKeyDescription = "description"
)

type SelectMsg struct {
	ActiveTransaction string
}

type Model struct {
	table       table.Model
	TableWidth  int
	TableMargin int
	db          *internal.JSONDatabase
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
			cmd = selectTransactionCmd(fmt.Sprintf("%v", selected))
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
	return constants.HelpStyle.Render(footer)
}

func (m *Model) loadTransactions() []models.TransactionModel {
	tempTransactions := m.db.Data.Transactions
	log.Printf("Found %d transactions\n", len(tempTransactions))
	transactionModels := make([]models.TransactionModel, len(tempTransactions))
	for i, transaction := range tempTransactions {
		transactionModels[i] = models.NewTransactionModel(transaction)
	}
	return transactionModels
}

func transformTransactionModelsToTableRows(transactionModels []models.TransactionModel) []table.Row {
	rows := make([]table.Row, len(transactionModels))
	for i, t := range transactionModels {
		var style lipgloss.Style
		if t.Transaction.Direction == internal.MoneyOutflow {
			style = lipgloss.NewStyle().Foreground(lipgloss.Color("#d11141"))
		} else {
			style = lipgloss.NewStyle().Foreground(lipgloss.Color("#00b159"))
		}

		rows[i] = table.NewRow(table.RowData{
			columnKeyValue:       style.Render(fmt.Sprintf("%.2f", t.Transaction.Value)),
			columnKeyTime:        t.Transaction.Time.Format("02-01-2006"),
			columnKeyPayee:       t.Transaction.Payee,
			columnKeyAccount:     t.Transaction.Account,
			columnKeyBudget:      t.Transaction.Budget,
			columnKeyDescription: t.Transaction.Description,
			columnKeyID:          t.Transaction.ID.String(),
		})
	}
	return rows
}

func createTableColumns() []table.Column {
	return []table.Column{
		table.NewFlexColumn(columnKeyValue, "Value", 8),
		table.NewColumn(columnKeyTime, "Time", 10),
		table.NewFlexColumn(columnKeyPayee, "Payee", 15),
		table.NewFlexColumn(columnKeyAccount, "Account", 15),
		table.NewFlexColumn(columnKeyBudget, "Budget", 15),
		table.NewFlexColumn(columnKeyDescription, "Description", 15),
	}
}

func NewOverviewModel(db *internal.JSONDatabase) Model {
	m := Model{db: db}
	tableRows := transformTransactionModelsToTableRows(m.loadTransactions())
	tableColumns := createTableColumns()

	keys := table.DefaultKeyMap()
	keys.RowUp.SetKeys(constants.Keymap.Up.Keys()...)
	keys.RowDown.SetKeys(constants.Keymap.Down.Keys()...)

	transactionTable := table.New(tableColumns).
		WithRows(tableRows).
		WithTargetWidth(60).
		WithKeyMap(keys).
		HeaderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#2a9d8f")).Bold(true)).
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
	m.table = transactionTable

	return m
}
