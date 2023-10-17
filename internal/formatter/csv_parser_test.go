package formatter

import (
	"strings"
	"testing"
	"time"

	"github.com/christhianjesus/storicard-challenge/internal/summarize"
	"github.com/stretchr/testify/require"
)

func TestEmptyGetTransactionsFromCSV(t *testing.T) {
	transactions, err := GetTransactionsFromCSV(strings.NewReader(``))

	require.NoError(t, err)
	require.Len(t, transactions, 0)
}

func TestEmptyWithHeaderGetTransactionsFromCSV(t *testing.T) {
	in := `Id,Date,Transaction`
	transactions, err := GetTransactionsFromCSV(strings.NewReader(in))

	require.NoError(t, err)
	require.Len(t, transactions, 0)
}

func TestParseErrorGetTransactionsFromCSV(t *testing.T) {
	transactions, err := GetTransactionsFromCSV(strings.NewReader(` 
	"""`))

	require.Error(t, err)
	require.Len(t, transactions, 0)
}

func TestLenErrorGetTransactionsFromCSV(t *testing.T) {
	transactions, err := GetTransactionsFromCSV(strings.NewReader(` 
	`))

	require.Error(t, err)
	require.Len(t, transactions, 0)
}

func TestGetTransactionsFromCSV(t *testing.T) {
	transactions, err := GetTransactionsFromCSV(strings.NewReader(`Id,Date,Transaction
0,7/15,+60.05`))

	require.NoError(t, err)
	require.Len(t, transactions, 1)

	transaction := transactions[0]
	require.Equal(t, 0, transaction.ID)
	require.Equal(t, time.July, transaction.Month)
	require.Equal(t, 60.05, transaction.Amount)
	require.Equal(t, summarize.DEBIT, transaction.Type)
}

func TestNilCSVRecordToTransaction(t *testing.T) {
	transactions, err := csvRecordToTransaction(nil)

	require.Error(t, err)
	require.Nil(t, transactions)
}

func TestIDErrorCSVRecordToTransaction(t *testing.T) {
	transactions, err := csvRecordToTransaction([]string{"", "", ""})

	require.Error(t, err)
	require.Nil(t, transactions)
}

func TestDateErrorCSVRecordToTransaction(t *testing.T) {
	transactions, err := csvRecordToTransaction([]string{"0", "", ""})

	require.Error(t, err)
	require.Nil(t, transactions)
}

func TestTransactionLenErrorCSVRecordToTransaction(t *testing.T) {
	transactions, err := csvRecordToTransaction([]string{"0", "0", ""})

	require.Error(t, err)
	require.Nil(t, transactions)
}

func TestSymbolErrorCSVRecordToTransaction(t *testing.T) {
	transactions, err := csvRecordToTransaction([]string{"0", "0", " "})

	require.Error(t, err)
	require.Nil(t, transactions)
}

func TestAmountErrorCSVRecordToTransaction(t *testing.T) {
	transactions, err := csvRecordToTransaction([]string{"0", "0", "-"})

	require.Error(t, err)
	require.Nil(t, transactions)
}
