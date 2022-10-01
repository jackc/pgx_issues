package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	var timeStamp time.Time
	err = conn.QueryRow(ctx, "select null::timestamptz").Scan(&timeStamp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(timeStamp)
}
