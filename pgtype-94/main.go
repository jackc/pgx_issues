package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	goUTCTime := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	var tstzPGBinary, tstzText time.Time

	err = conn.QueryRow(
		context.Background(),
		`select '2010-01-01 00:00:00+00'::timestamptz, '2010-01-01 00:00:00+00'::timestamptz`,
		pgx.QueryResultFormats{1, 0},
	).Scan(&tstzPGBinary, &tstzText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("goUTCTime: %#v\n", goUTCTime)
	fmt.Printf("tstzPGBinary: %#v\n", tstzPGBinary)
	fmt.Printf("tstzText: %#v\n", tstzText)
}
