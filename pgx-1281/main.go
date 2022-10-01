package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
)

func main() {
	conn, err := pgconn.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	rr := conn.ExecParams(context.Background(), `select 1 as foofoofoo, 2 as barbarbar, 3 as bazbazbaz, 4 as quzquzquz from generate_series(1, 5000) n`, nil, nil, nil, nil)
	for rr.NextRow() {
		for _, fd := range rr.FieldDescriptions() {
			fmt.Println(string(fd.Name))
		}
	}

	_, err = rr.Close()
	if err != nil {
		log.Fatal(err)
	}
}
