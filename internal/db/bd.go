package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(databaseURL string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Erro ao parsear config do DB:", err)
	}

	// Configurações do pool
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	DB, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("Erro ao conectar ao DB:", err)
	}

	// Testa conexão
	err = DB.Ping(ctx)
	if err != nil {
		log.Fatal("DB não responde:", err)
	}

	log.Println("Conectado ao banco")
}
