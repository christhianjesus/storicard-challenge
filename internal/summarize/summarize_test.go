package summarize

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEmptySummarize(t *testing.T) {
	summary := Summarize(TransactionList{})

	require.Equal(t, 0.0, summary.TotalBalance)
	require.Len(t, summary.SummaryPerMonth, 0)
}

func TestSummarize(t *testing.T) {
	summary := Summarize(getTestTransactionList())

	require.Equal(t, 7.5, summary.TotalBalance)
	require.Len(t, summary.SummaryPerMonth, 2)

	april := summary.SummaryPerMonth[0]
	require.Equal(t, time.April.String(), april.MonthName)
	require.Equal(t, 1, april.TransactionsQTY)
	require.Equal(t, 3.5, april.AVGDebit)
	require.Equal(t, 0.0, april.AVGCredit)

	june := summary.SummaryPerMonth[1]
	require.Equal(t, time.June.String(), june.MonthName)
	require.Equal(t, 2, june.TransactionsQTY)
	require.Equal(t, 1.5, june.AVGDebit)
	require.Equal(t, 2.5, june.AVGCredit)
}
