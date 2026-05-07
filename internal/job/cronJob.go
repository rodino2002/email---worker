package job

import (
	"email-WORKER/internal/service"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func StartMessagesGOSMSJob() {
	log.Print("Configurando o CRONJOB Para verificar mensagens pendentes")
	loc, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		log.Fatalf("Erro ao carregar timezone: %v", err)
	}

	c := cron.New(cron.WithSeconds(), cron.WithLocation(loc))
	//	0 */5 * * * * para 5minutos
	//	@every 5s: */5 * * * * * para teste rápido
	_, err =
		c.AddFunc("0 */30 * * * *", func() {
			service.CheckAndNotify()
		})
	c.Start()
	log.Print("CRONJOB configurado com sucesso")

}
