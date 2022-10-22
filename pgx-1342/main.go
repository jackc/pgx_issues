package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	poolConfig.ConnConfig.RuntimeParams["statement_timeout"] = "5000"
	poolConfig.ConnConfig.RuntimeParams["idle_in_transaction_session_timeout"] = "60000"

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	var statementTimeout string
	err = pool.QueryRow(ctx, "show statement_timeout").Scan(&statementTimeout)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("statement_timeout", statementTimeout)

	var idleInTransactionSessionTimeout string
	err = pool.QueryRow(ctx, "show idle_in_transaction_session_timeout").Scan(&idleInTransactionSessionTimeout)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("idle_in_transaction_session_timeout", idleInTransactionSessionTimeout)
}
