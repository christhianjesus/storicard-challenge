package summarize

import "time"

type TransactionType string

const (
	DEBIT  TransactionType = "Debit"
	CREDIT TransactionType = "Credit"
)

type Transaction struct {
	ID     int
	Month  time.Month
	Amount float64
	Type   TransactionType
}
