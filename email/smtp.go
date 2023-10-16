package email

import (
	"context"
	"net/smtp"
	"strings"
)

type MailConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

type MailContent struct {
	From    string
	To      []string
	Subject string
	Body    string
}

type EmailSender interface {
	SendEmail(ctx context.Context, content *MailContent) error
}

func NewHTMLEmailSender(config *MailConfig) EmailSender {
	addr := config.Host + ":" + config.Port
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	return &htmlEmailSender{
		addr:     addr,
		auth:     auth,
		sendMail: smtp.SendMail,
	}
}

type sendMailFunc = func(string, smtp.Auth, string, []string, []byte) error

type htmlEmailSender struct {
	addr     string
	auth     smtp.Auth
	sendMail sendMailFunc
}

func (hes *htmlEmailSender) SendEmail(ctx context.Context, content *MailContent) error {
	msg := buildHTMLMessage(content)

	return hes.sendMail(hes.addr, hes.auth, content.From, content.To, msg)
}

func buildHTMLMessage(content *MailContent) []byte {
	msg := []string{
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=utf-8",
		"From: " + content.From,
		"To: " + strings.Join(content.To, ";"),
		"Subject: " + content.Subject,
		"",
		content.Body,
	}

	return []byte(strings.Join(msg, "\r\n"))
}
