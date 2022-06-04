package models

import (
	"github.com/rwxd/zaster/internal"
)

type TransactionModel struct {
	Transaction internal.Transaction
}

func (t TransactionModel) Title() string {
	return t.Transaction.Str()
}

func (t TransactionModel) Description() string {
	return t.Transaction.Description
}

func (t TransactionModel) FilterValue() string {
	return t.Transaction.Str()
}

func NewTransactionModel(transaction internal.Transaction) TransactionModel {
	return TransactionModel{Transaction: transaction}
}
