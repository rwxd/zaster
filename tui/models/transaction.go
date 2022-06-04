package models

import (
	"github.com/rwxd/zaster/internal"
)

type TransactionModel struct {
	transaction internal.Transaction
}

func (t TransactionModel) Title() string {
	return t.transaction.Str()
}

func (t TransactionModel) Description() string {
	return t.transaction.Description
}

func (t TransactionModel) FilterValue() string {
	return t.transaction.Str()
}

func NewTransactionModel(transaction internal.Transaction) TransactionModel {
	return TransactionModel{transaction: transaction}
}
