package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()

	if err := run(ctx, db); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, db *sql.DB) error {
	txOptions := &sql.TxOptions{ReadOnly: true}
	for {
		txn, err := db.BeginTx(ctx, txOptions)
		if err != nil {
			return fmt.Errorf("error starting transaction: %w", err)
		}

		if _, err := txn.Exec(
			"SET LOCAL statement_timeout TO 1",
		); err != nil {
			_ = txn.Rollback()
			return fmt.Errorf("error setting statement_timeout: %w", err)
		}

		row := txn.QueryRow("SELECT 1, pg_sleep(1)")
		var id int
		if err := row.Scan(&id, nil); err != nil {
			log.Printf("error scanning row: %s", err)
			// _ = txn.Rollback()
			// continue
		}

		if err := txn.Commit(); err != nil {
			log.Printf("error committing: %s", err)
			_ = txn.Rollback()
			continue
		}
	}
}
