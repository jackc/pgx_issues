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

	_, err = conn.Exec(context.Background(), `drop table if exists basic_types;

	CREATE TABLE basic_types
(
    id            BIGSERIAL PRIMARY KEY,
    small_int     INT2,
    default_int   INT,
    medium_int    INT4,
    large_int     INT8,
    regular_char  CHAR
);

INSERT INTO basic_types (small_int, default_int, medium_int, large_int, regular_char)
VALUES (5, 5, 5, 5, 'A');

INSERT INTO basic_types DEFAULT VALUES ;
	`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable setup test data: %v\n", err)
		os.Exit(1)
	}

	var smallInt *int16
	var defaultInt *int32
	var mediumInt *int32
	var largeInt *int64
	var regularChar *rune

	rows, err := conn.Query(context.Background(), `select small_int, default_int, medium_int, large_int, regular_char from basic_types order by id asc`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	if !rows.Next() {
		fmt.Fprintln(os.Stderr, "Expected next row but none received")
		os.Exit(1)
	}

	fmt.Println("before first scan")

	err = rows.Scan(&smallInt, &defaultInt, &mediumInt, &largeInt, &regularChar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scan of non-nulls failed: %v\n", err)
		os.Exit(1)
	}

	if !rows.Next() {
		fmt.Fprintln(os.Stderr, "Expected next row but none received")
		os.Exit(1)
	}

	fmt.Println("before second scan")

	err = rows.Scan(&smallInt, &defaultInt, &mediumInt, &largeInt, &regularChar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scan of nulls failed: %v\n", err)
		os.Exit(1)
	}

	if regularChar != nil {
		fmt.Fprintf(os.Stderr, "Expected regularChar to be nil but it is: %v\n", regularChar)
		os.Exit(1)
	}

	if rows.Next() {
		fmt.Fprintln(os.Stderr, "Unexpected next row")
		os.Exit(1)
	}
	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows.Err(): %v\n", rows.Err())
		os.Exit(1)
	}
}
