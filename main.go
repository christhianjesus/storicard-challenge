package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/christhianjesus/storicard-challenge/internal/email"
	"github.com/christhianjesus/storicard-challenge/internal/formatter"
	"github.com/christhianjesus/storicard-challenge/internal/summarize"
)

var (
	svc          *s3.Client
	templateName string
	emailSender  email.EmailSender
	emailContent *email.MailContent
)

func init() {
	// Load the SDK configuration
	region := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	// Initialize an S3 client
	svc = s3.NewFromConfig(cfg)

	// Get template name
	templateName = os.Getenv("TEMPLATE_NAME")

	// Load SMTP credentials
	emailSender = email.NewHTMLEmailSender(&email.MailConfig{
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	})

	// Email fixed content
	emailContent = &email.MailContent{
		Subject: os.Getenv("EMAIL_SUBJECT"),
		From:    os.Getenv("EMAIL_FROM"),
	}
}

type MyEvent struct {
	Body string `json:"body"`
}

type Body struct {
	EmailTo  string `json:"email_to"`
	FileName string `json:"csv_file_name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) error {
	// Unmarshall body
	var body Body
	if err := json.Unmarshal([]byte(event.Body), &body); err != nil {
		return fmt.Errorf("Error parsing json: %w", err)
	}

	// Read event
	emailContent.To = strings.Split(body.EmailTo, ";")
	input := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(body.FileName),
	}

	// Get file
	file, err := svc.GetObject(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Body.Close()

	// Parse .csv
	transactions, err := formatter.GetTransactionsFromCSV(file.Body)
	if err != nil {
		return fmt.Errorf("Error reading csv: %w", err)
	}

	// Summarize transactions
	summary := summarize.Summarize(transactions)

	// Render template
	emailContent.Body, err = formatter.GenerateTransactionEmail(templateName, summary)
	if err != nil {
		return fmt.Errorf("Error rendering template: %w", err)
	}

	// Send email
	if err := emailSender.SendEmail(context.TODO(), emailContent); err != nil {
		return fmt.Errorf("Error sending email: %w", err)
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
