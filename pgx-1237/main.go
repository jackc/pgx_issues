package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, `host=localhost user=jack password="something\""`)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	var n int32
	err = conn.QueryRow(ctx, "select 1").Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(n)
}
