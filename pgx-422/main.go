package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `create or replace procedure testproc(inout foo int) language plpgsql as $$
	begin
		foo := 42;
	end;
	$$;`)
	if err != nil {
		log.Fatal(err)
	}

	var foo int
	err = conn.QueryRow(context.Background(), "call testproc(1)").Scan(&foo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("foo", foo)
}
