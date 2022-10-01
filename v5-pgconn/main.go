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

	rr := conn.ExecParams(context.Background(), `select 1`, nil, nil, nil, nil)
	for rr.NextRow() {
		fmt.Println(rr.Values())
	}

	_, err = rr.Close()
	if err != nil {
		log.Fatal(err)
	}
}
