package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	n1 := &pgtype.Numeric{NaN: true, Status: pgtype.Present}
	n2 := &pgtype.Numeric{}

	err = conn.QueryRow(context.Background(), "select $1::numeric", n1).Scan(n2)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(n1)
	fmt.Println(n2)
}
