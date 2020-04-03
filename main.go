package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())
}
