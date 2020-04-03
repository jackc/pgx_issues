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
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `drop type if exists custom; create type custom as (a int, b text);`)
	if err != nil {
		log.Fatalln(err)
	}

	genericBinary := &pgtype.GenericBinary{}
	err = conn.QueryRow(context.Background(), `select '(42,foo)'::custom`, pgx.QueryResultFormats{pgx.BinaryFormatCode}).Scan(genericBinary)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(genericBinary.Bytes)

}
