package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.TODO()

	pgxCfg, err := pgxpool.ParseConfig(os.Getenv("TEST_POSTGRES"))
	if err != nil {
		panic(fmt.Errorf("could not parse DB URL: %w", err))
	}

	pgxPool, err := pgxpool.ConnectConfig(ctx, pgxCfg)
	if err != nil {
		panic(fmt.Errorf("could not init pgx conn pool: %w", err))
	}

	if err := pgxPool.Ping(ctx); err != nil {
		panic(fmt.Errorf("could not ping DB: %w", err))
	}

	var arr pgtype.Int8Array
	if err := pgxPool.QueryRow(ctx, `SELECT ARRAY[1, 2]::bigint[]`).Scan(&arr); err != nil {
		panic(fmt.Errorf("could not read static array value: %w", err))
	}
}
