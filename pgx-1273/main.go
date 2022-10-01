package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type Example struct {
	id        int
	stub      bool
	timestamp time.Time
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	var example Example
	err = conn.QueryRow(ctx, "select id, data->'stub' from example where id=$1", 1).Scan(&example.id, &example.stub)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(example)
}
