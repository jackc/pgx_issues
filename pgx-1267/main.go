package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "truncate table public.testtable")
	if err != nil {
		log.Fatal(err)
	}

	inputRows := [][]interface{}{
		{1, &pgtype.Box{P: [2]pgtype.Vec2{{X: 0, Y: 0}, {X: 100, Y: 100}}, Status: pgtype.Present}},
		{2, &pgtype.Box{P: [2]pgtype.Vec2{{X: 0, Y: 0}, {X: 200, Y: 200}}, Status: pgtype.Present}},
	}
	copyCount, err := conn.CopyFrom(context.Background(), pgx.Identifier{"public", "testtable"}, []string{"id", "fbox"}, pgx.CopyFromRows(inputRows))
	log.Println(copyCount, err)
}
