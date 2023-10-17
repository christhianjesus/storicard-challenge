package email

import (
	"context"
	"errors"
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTMLEmailSender(t *testing.T) {
	config := &MailConfig{
		Username: "username",
		Password: "password",
		Host:     "host",
		Port:     "port",
	}
	emailSender := NewHTMLEmailSender(config)

	require.NotNil(t, emailSender)
	require.IsType(t, &htmlEmailSender{}, emailSender)

	htmlEmailSenderConcrete := emailSender.(*htmlEmailSender)

	assert.Equal(t, "host:port", htmlEmailSenderConcrete.addr)
	assert.NotNil(t, htmlEmailSenderConcrete.auth)
	assert.NotNil(t, htmlEmailSenderConcrete.sendMail)
}

func TestErrorSendEmail(t *testing.T) {
	ef := func(string, smtp.Auth, string, []string, []byte) error {
		return errors.New("")
	}
	mailSender := &htmlEmailSender{
		sendMail: ef,
	}
	err := mailSender.SendEmail(context.TODO(), &MailContent{})

	require.Error(t, err)
}

func TestNoErrorSendEmail(t *testing.T) {
	nef := func(string, smtp.Auth, string, []string, []byte) error {
		return nil
	}
	mailSender := &htmlEmailSender{
		sendMail: nef,
	}
	err := mailSender.SendEmail(context.TODO(), &MailContent{})

	require.NoError(t, err)
}

func TestBuildHTMLMessage(t *testing.T) {
	msg := string(buildHTMLMessage(&MailContent{
		From:    "from@email.com",
		To:      []string{"to@email.com"},
		Subject: "email title",
		Body:    "email body",
	}))

	require.Contains(t, msg, "MIME-Version: 1.0")
	require.Contains(t, msg, "Content-Type: text/html; charset=utf-8")
	require.Contains(t, msg, "From: from@email.com")
	require.Contains(t, msg, "To: to@email.com")
	require.Contains(t, msg, "Subject: email title")
	require.Contains(t, msg, "email body")
}
