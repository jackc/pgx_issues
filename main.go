package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgproto3"
)

func main() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())

	conn.PgConn().Frontend().MessageTracer = &pgproto3.LibpqMessageTracer{
		Writer: os.Stderr,
	}

	rows, _ := conn.Query(context.Background(), "select * from generate_series(1, 10)")
	rows.Close()
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}
}
