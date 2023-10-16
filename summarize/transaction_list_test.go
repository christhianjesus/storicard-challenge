package summarize

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getTestTransactionList() TransactionList {
	return TransactionList{
		&Transaction{
			ID:     1,
			Month:  time.April,
			Amount: 3.5,
			Type:   DEBIT,
		},
		&Transaction{
			ID:     2,
			Month:  time.June,
			Amount: 2.5,
			Type:   CREDIT,
		},
		&Transaction{
			ID:     3,
			Month:  time.June,
			Amount: 1.5,
			Type:   DEBIT,
		},
	}
}

func TestEmptyTotalBalance(t *testing.T) {
	total := TransactionList{}.TotalBalance()

	require.Equal(t, 0.0, total)
}

func TestTotalBalance(t *testing.T) {
	total := getTestTransactionList().TotalBalance()

	require.Equal(t, 7.5, total)
}

func TestEmptyAverage(t *testing.T) {
	avg := TransactionList{}.Average("")

	require.Equal(t, 0.0, avg)
}

func TestAverageInvalidType(t *testing.T) {
	avg := getTestTransactionList().Average("")

	require.Equal(t, 0.0, avg)
}

func TestAverageDebit(t *testing.T) {
	avg := getTestTransactionList().Average(DEBIT)

	require.Equal(t, 2.5, avg)
}

func TestAverageCredit(t *testing.T) {
	avg := getTestTransactionList().Average(CREDIT)

	require.Equal(t, 2.5, avg)
}

func TestEmptyGroupByMonth(t *testing.T) {
	groups := TransactionList{}.GroupByMonth()

	require.Len(t, groups, 0)
}

func TestGroupByMonth(t *testing.T) {
	groups := getTestTransactionList().GroupByMonth()

	require.Len(t, groups, 2)
	require.Len(t, groups[time.April], 1)
	require.Len(t, groups[time.June], 2)
	require.Equal(t, 1, groups[time.April][0].ID)
	require.Equal(t, 2, groups[time.June][0].ID)
	require.Equal(t, 3, groups[time.June][1].ID)
}
