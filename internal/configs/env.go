package configs

import (
	"email-WORKER/internal/errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type VariableEnv struct {
	DatabaseURL string
	Port        string
}

func GetEnvVariables() (*VariableEnv, error) {
	var dataBase_url string
	if os.Getenv("APP_ENV") != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			godotenv.Load()
		}
		dataBase_url = os.Getenv("DATABASE_URL_DEV")
		fmt.Println("desenvolvimento")
	} else {
		dataBase_url = os.Getenv("DATABASE_URL_PROD")
		fmt.Println("producao")

	}

	if dataBase_url == "" {
		return nil, errors.ErrMissingDataBaseUrl
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &VariableEnv{
		DatabaseURL: dataBase_url,
		Port:        port,
	}, nil

}
