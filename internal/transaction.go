package internal

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MoneyDirection string

const MoneyInflow MoneyDirection = "in"
const MoneyOutflow MoneyDirection = "out"

type Transaction struct {
	ID          uuid.UUID      `json:"id"`
	Value       float64        `json:"value"`
	Time        time.Time      `json:"date"`
	Payee       string         `json:"payee"`
	Account     Account        `json:"account"`
	Budget      Budget         `json:"budget"`
	Description string         `json:"description"`
	Direction   MoneyDirection `json:"direction"`
}

func (t Transaction) Str() string {
	var prefix string
	if t.Direction == MoneyInflow {
		prefix = "From"
	} else {
		prefix = "To"
	}
	return fmt.Sprintf("%s %s, from account %s, amount %f, on %s", prefix, t.Payee, t.Account, t.Value, t.Time.Format("02-01-2006"))
}

func NewTransaction(
	value float64,
	time time.Time,
	payee string,
	account Account,
	budget Budget,
	description string,
	direction MoneyDirection,
) (Transaction, error) {
	transaction := Transaction{
		ID:          uuid.New(),
		Value:       value,
		Time:        time,
		Payee:       payee,
		Account:     account,
		Description: description,
		Budget:      budget,
		Direction:   direction,
	}
	return transaction, nil
}

// func GetFilteredTransactions(transactions []Transaction, date time.Time, receiver string,
// 	account string, category string, budget string, description string) ([]Transaction, error) {
//
// 	var filtered []Transaction
//
// 	for _, transaction := range transactions {
// 		if date == transaction.Time {
// 			filtered = append(filtered, transaction)
// 			break
// 		} else if receiver != "" && receiver == transaction.Payee {
// 			filtered = append(filtered, transaction)
// 			break
// 		} else if category != "" && category == transaction.Category {
// 			filtered = append(filtered, transaction)
// 			break
// 		} else if account != "" && account == transaction.Account {
// 			filtered = append(filtered, transaction)
// 			break
// 		} else if description != "" && strings.Contains(description, transaction.Description) {
// 			filtered = append(filtered, transaction)
// 			break
// 		} else if budget != "" && budget == transaction.Budget {
// 			filtered = append(filtered, transaction)
// 			break
// 		}
// 	}
// 	return filtered, nil
// }
