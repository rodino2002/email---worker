package service

import (
	"email-WORKER/internal/db"
	"email-WORKER/internal/mailer"
	"fmt"
	"os"
)

// CheckAndNotify valida estado do sistema de SMS (health monitor)
func CheckAndNotify() error {

	// 1. mensagens pendentes na fila
	pending, err := db.CountPendingMessages()
	if err != nil {
		return err
	}

	// 2. mensagens processadas recentemente (health check)
	processed, err := db.CountProcessedMessages()
	if err != nil {
		return err
	}

	fmt.Printf("SMS pendentes: %d\n", pending)
	fmt.Printf("SMS processadas (5min): %d\n", processed)

	// 3. sistema sem atividade nenhuma
	if pending == 0 && processed == 0 {
		fmt.Println("Sem tráfego no sistema")
		return nil
	}

	// 4. sistema saudável (está processando normalmente)
	if processed > 0 {
		fmt.Println("Sistema funcionando normalmente")
		return nil
	}

	// 5. caso crítico: há fila mas nada está a ser processado
	if pending > 0 && processed == 0 {
		fmt.Println(" Sistema possivelmente parado (sem processamento recente)")

		recipientEmail := os.Getenv("EMAIL_TO")

		err = mailer.SendEmail(mailer.EmailData{
			To:      recipientEmail,
			Subject: " ALERTA - GOSMS System Down",
			Body: fmt.Sprintf(
				"ALERTA!\n\nO sistema possui %d SMS pendentes,\nmas nenhuma foi processada nos últimos 5 minutos.\n\nPossível falha no envio de SMS.",
				pending,
			),
		})

		if err != nil {
			return err
		}

		fmt.Println("Email de alerta enviado com sucesso!")
		return nil
	}

	return nil
}
