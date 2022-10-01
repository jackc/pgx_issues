package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var id *uuid.UUID
	err = conn.QueryRow(context.Background(), "select gen_random_uuid()").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Random UUID", id)

	err = conn.QueryRow(context.Background(), "select null::uuid").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("NULL UUID", id)
}
