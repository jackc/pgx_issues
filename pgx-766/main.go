package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, "create temporary table pgx766(a text, b text, c text)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create temporary table: %v\n", err)
		os.Exit(1)
	}

	rows, err := conn.Query(ctx, "select * from pgx766")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(len(rows.FieldDescriptions()), "columns")
	for _, fd := range rows.FieldDescriptions() {
		fmt.Println(string(fd.Name))
	}

	rows.Close()
	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows.Err(): %v\n", rows.Err())
		os.Exit(1)
	}

}
