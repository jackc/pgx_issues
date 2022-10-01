package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var dst *[]string
	err = conn.QueryRow(context.Background(), `SELECT ARRAY['foo val', 'foo val 2', 'foo val 3']`).Scan(&dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to select pointer to slice: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(dst)
}
