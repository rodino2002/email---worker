package mailer

import (
	"fmt"
	"net/smtp"
	"os"
)

type EmailData struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(data EmailData) error {
	from := os.Getenv("EMAIL_FROM")
	pass := os.Getenv("EMAIL_PASS")

	if from == "" || pass == "" {
		return fmt.Errorf("credenciais de email não definidas")
	}

	msg := buildMessage(from, data)

	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		[]string{data.To},
		[]byte(msg),
	)

	if err != nil {
		return err
	}

	return nil
}

func buildMessage(from string, data EmailData) string {
	return fmt.Sprintf(
		"From: %s\nTo: %s\nSubject: %s\n\n%s",
		from,
		data.To,
		data.Subject,
		data.Body,
	)
}
