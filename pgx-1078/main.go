package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `CREATE EXTENSION IF NOT EXISTS citext;
	DROP TABLE IF EXISTS test;
	CREATE TABLE test (
			mytext citext
	);
	INSERT INTO test (mytext)
		VALUES (NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	results, err := Test(context.Background(), conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(results)
}

const test = `-- name: Test :many
SELECT
  mytext
FROM
  test
`

func Test(ctx context.Context, conn *pgx.Conn) ([]sql.NullString, error) {
	rows, err := conn.Query(ctx, test)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []sql.NullString
	for rows.Next() {
		var mytext sql.NullString
		if err := rows.Scan(&mytext); err != nil {
			return nil, err
		}
		items = append(items, mytext)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
