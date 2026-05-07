package main

import (
	"email-WORKER/internal/configs"
	"email-WORKER/internal/db"
	"email-WORKER/internal/job"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Iniciando worker...")

	envs, err := configs.GetEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	db.Connect(envs.DatabaseURL)
	job.StartMessagesGOSMSJob()

	// route := gin.Default()

	// route.Run(":8081") // listen and serve on
	select {}

}
