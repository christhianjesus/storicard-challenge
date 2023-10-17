package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/christhianjesus/storicard-challenge/internal/email"
	"github.com/christhianjesus/storicard-challenge/internal/formatter"
	"github.com/christhianjesus/storicard-challenge/internal/summarize"
)

func main() {
	// SMTP credentials
	emailSender := email.NewHTMLEmailSender(&email.MailConfig{
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	})

	// Email data
	fileName := os.Getenv("CSV_FILE_NAME")
	templateName := os.Getenv("TEMPLATE_NAME")
	content := &email.MailContent{
		From:    os.Getenv("EMAIL_FROM"),
		To:      strings.Split(os.Getenv("EMAIL_TO"), ";"),
		Subject: os.Getenv("EMAIL_SUBJECT"),
	}

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Parse .csv
	transactions, err := formatter.GetTransactionsFromCSV(file)
	if err != nil {
		log.Fatalf("Error reading csv: %v", err)
	}

	// Summarize transactions
	summary := summarize.Summarize(transactions)

	// Render template
	content.Body, err = formatter.GenerateTransactionEmail(templateName, summary)
	if err != nil {
		log.Fatalf("Error rendering template: %v", err)
	}

	// Send email
	if err := emailSender.SendEmail(context.TODO(), content); err != nil {
		log.Fatalf("Error sending email: %v", err)
	}
}
