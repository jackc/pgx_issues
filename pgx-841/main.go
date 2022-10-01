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

	_, err = conn.Prepare(context.Background(), "s1", "select 1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to prepare statement: %v\n", err)
		os.Exit(1)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to begin transaction: %v\n", err)
		os.Exit(1)
	}
	defer tx.Rollback(context.Background())

	var f float64
	err = conn.QueryRow(context.Background(), "select 1 / 0").Scan(&f)
	if err == nil {
		fmt.Fprintln(os.Stderr, "Expected error but did not receive")
		os.Exit(1)
	}

	err = conn.Deallocate(context.Background(), "s1")
	fmt.Println("Deallocate in broken transaction error:", err)
}
