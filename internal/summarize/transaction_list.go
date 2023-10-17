package summarize

import "time"

type TransactionList []*Transaction

func (t TransactionList) TotalBalance() float64 {
	var totalBalance float64

	for _, e := range t {
		totalBalance += e.Amount
	}

	return totalBalance
}

func (t TransactionList) Average(transactionType TransactionType) float64 {
	var average float64

	quantity := 0
	for _, e := range t {
		if e.Type == transactionType {
			average += e.Amount
			quantity += 1
		}
	}

	if quantity > 0 {
		return average / float64(quantity)
	}

	return 0.0
}

func (t TransactionList) GroupByMonth() map[time.Month]TransactionList {
	groups := make(map[time.Month]TransactionList, 12)

	for _, e := range t {
		groups[e.Month] = append(groups[e.Month], e)
	}

	return groups
}
