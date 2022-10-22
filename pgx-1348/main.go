package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	DB_URL := os.Getenv("DATABASE_URL")
	panicOnErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	conn, err := pgx.Connect(ctx, DB_URL)
	panicOnErr(err)
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, `
	  CREATE TEMPORARY TABLE foo (
		stuff int[]
	  )
	`)
	panicOnErr(err)

	qry := `
	  INSERT INTO foo (stuff) values ($1)
	`
	concreteRows := []int{1, 2, 3, 4, 5}
	_, err = conn.Exec(ctx, qry, concreteRows) // Works fine
	panicOnErr(err)

	interfaceRows := []interface{}{1, 2, 3, 4, 5} // SIGSEGV!
	_, err = conn.Exec(ctx, qry, interfaceRows)
	panicOnErr(err)
}
