package main

import (
	"context"
	"fmt"
	"log"
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

	// select two rows: (NULL, NULL) and (1, 2)
	rows, err := conn.Query(context.Background(), `
		SELECT NULL as a, NULL as b
		UNION
		SELECT 1, 2
	`)

	if !rows.Next() {
		log.Fatal("no next")
	}

	// scan nil: ignore the first row
	if err := rows.Scan(nil, nil); err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		log.Fatal("no next")
	}

	var a int
	var b int
	if err := rows.Scan(&a, &b); err != nil {
		log.Fatal(err)
	}

	fmt.Println(a, b)
}
