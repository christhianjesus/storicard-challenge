package formatter

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/christhianjesus/storicard-challenge/summarize"
)

func GetTransactionsFromCSV(file io.Reader) (summarize.TransactionList, error) {
	reader := csv.NewReader(file)

	// Ignore header
	reader.Read()

	var transactions summarize.TransactionList
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		transaction, err := csvRecordToTransaction(record)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func csvRecordToTransaction(record []string) (*summarize.Transaction, error) {
	if len(record) != 3 {
		return nil, errors.New("Invalid line len")
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, errors.New("Invalid 'id' field")
	}

	date := strings.Split(record[1], "/")
	month, err := strconv.Atoi(date[0])
	if err != nil {
		return nil, errors.New("Invalid 'date' field")
	}

	if len(record[2]) == 0 {
		return nil, errors.New("Invalid 'transaction' field")
	}

	var transactionType summarize.TransactionType
	switch symbol := record[2][:1]; symbol {
	case "+":
		transactionType = summarize.DEBIT
	case "-":
		transactionType = summarize.CREDIT
	default:
		return nil, errors.New("Invalid 'transaction' field")
	}

	amount, err := strconv.ParseFloat(record[2][1:], 64)
	if err != nil {
		return nil, errors.New("Invalid 'transaction' field")
	}

	return &summarize.Transaction{
		ID:     id,
		Month:  time.Month(month),
		Amount: amount,
		Type:   transactionType,
	}, nil
}
