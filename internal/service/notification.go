package service

import (
	"email-WORKER/internal/db"
	"email-WORKER/internal/mailer"
	"fmt"
	"os"
)

const minimumPendingMessages = 100

func CheckAndNotify() error {

	count, err := db.CountPendingMessages()
	if err != nil {
		return err
	}

	fmt.Printf("SMS pendentes: %d\n", count)

	if count < minimumPendingMessages {
		fmt.Println("Quantidade insuficiente para notificação")
		return nil
	}

	recipientEmail := os.Getenv("EMAIL_TO")

	err = mailer.SendEmail(mailer.EmailData{
		To:      recipientEmail,
		Subject: "Sistema de Monitoramento - GOSMS",
		Body: fmt.Sprintf(
			"Olá, o sistema identificou mensagens pendentes na fila. \n Quantidade atual: %d ", count),
	})

	if err != nil {
		return err
	}

	fmt.Println("Email enviado com sucesso!")

	return nil
}
