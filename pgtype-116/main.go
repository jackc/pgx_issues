package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int64
	var changes pgtype.Hstore
	err = conn.QueryRow(context.Background(), `select * from test`, pgx.QueryResultFormats{pgx.BinaryFormatCode}).Scan(&id, &changes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable select hstore: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(changes)
}
