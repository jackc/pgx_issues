package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

const (
	query = `
UPDATE
    test_text
SET
    (name, num) = (vals.column2, vals.column3)
FROM
    (
        VALUES
            ($1::INT, $2::STRING, $3::INT),
            ($4::INT, $5::STRING, $6::INT)
    ) AS vals
WHERE
    id = vals.column1
`
)

func main() {
	connStr := ""

	err := UpdateWithPQ(connStr)
	if err != nil {
		fmt.Printf("PQ exec with error: %v\n", err)
	} else {
		fmt.Println("PQ processed successfully")
	}

	err = UpdateWithPGX("")
	if err != nil {
		fmt.Printf("PGX exec with error: %v\n", err)
	} else {
		fmt.Println("PGX processed successfully")
	}

	err = UpdateWithPGXString(connStr)
	if err != nil {
		fmt.Printf("PGXString exec with error: %v\n", err)
	} else {
		fmt.Println("PGXString processed successfully")
	}
}

func UpdateWithPQ(connStr string) error {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer conn.Close() // nolint: errcheck

	values := []interface{}{
		1, "one...", 101,
		2, "two...", 102,
	}

	_, err = conn.Exec(query, values...) // OK
	return err
}

func UpdateWithPGX(connStr string) error {
	ctx := context.Background()
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer conn.Close(ctx) // nolint: errcheck
	values := []interface{}{
		1, "one...", 102,
		2, "two...", 101,
	}

	_, err = conn.Exec(ctx, query, values...) // Failed
	return err
}

func UpdateWithPGXString(connStr string) error {
	ctx := context.Background()
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer conn.Close(ctx) // nolint: errcheck

	values := []interface{}{
		fmt.Sprintf("%d", 1), "one...", fmt.Sprintf("%d", 101),
		fmt.Sprintf("%d", 2), "two...", fmt.Sprintf("%d", 102),
	}
	_, err = conn.Exec(ctx, query, values...) // OK
	return err
}
