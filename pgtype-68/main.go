package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var a uuid.UUID
	err = conn.QueryRow(context.Background(), `select gen_random_uuid()`).Scan(&a)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable select uuid: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(a)

	var t []uuid.UUID
	err = conn.QueryRow(context.Background(), `select '{}'::uuid[]`).Scan(&t)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable select empty uuid array: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(t)
}
