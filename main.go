package main

import (
	"context"
	"os"

	"github.com/jackc/pgconn"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := pgconn.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	must(err)
	defer conn.Close(context.Background())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn.Exec(ctx, "SELECT table_name FROM information_schema.tables")
	conn.Close(ctx)
}
