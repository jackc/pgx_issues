package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	n := decimal.RequireFromString("1.234")

	err = conn.QueryRow(ctx, "select $1::numeric", n).Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(n)
}
