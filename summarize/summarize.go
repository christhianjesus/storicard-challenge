package summarize

import "time"

type MonthSummary struct {
	MonthName       string
	TransactionsQTY int
	AVGDebit        float64
	AVGCredit       float64
}

type Summary struct {
	TotalBalance    float64
	SummaryPerMonth []*MonthSummary
}

func Summarize(transactionList TransactionList) *Summary {
	transactionsPerMonth := transactionList.GroupByMonth()

	summaryPerMonth := make([]*MonthSummary, 0, len(transactionsPerMonth))
	for i := 1; i <= 12; i++ {
		month := time.Month(i)
		if transactions, ok := transactionsPerMonth[month]; ok {
			summaryPerMonth = append(summaryPerMonth, &MonthSummary{
				MonthName:       month.String(),
				TransactionsQTY: len(transactions),
				AVGDebit:        transactions.Average(DEBIT),
				AVGCredit:       transactions.Average(CREDIT),
			})
		}
	}

	return &Summary{
		TotalBalance:    transactionList.TotalBalance(),
		SummaryPerMonth: summaryPerMonth,
	}
}
