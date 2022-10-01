package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var s string
	err = conn.QueryRow(context.Background(), `select 1.234::numeric`).Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)

}
