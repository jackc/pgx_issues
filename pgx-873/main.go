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

	rows, _ := conn.Query(context.Background(), "select 'foo', 42")
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Values error: %v\n", err)
			os.Exit(1)
		}

		for _, v := range values {
			fmt.Printf("%T\n", v)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Query error: %v\n", err)
		os.Exit(1)
	}
}
