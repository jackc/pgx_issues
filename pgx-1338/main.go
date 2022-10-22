package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	panicMaybe(err)

	_, err = conn.Exec(
		context.Background(),
		`
		create type my_enum as enum('yes','no','toaster');

		create table example(
			value my_enum
		)`,
	)
	panicMaybe(err)

	t, err := conn.LoadType(context.Background(), "my_enum")
	panicMaybe(err)
	conn.TypeMap().RegisterType(t)

	_, err = conn.CopyFrom(context.Background(), []string{"example"}, []string{"value"}, pgx.CopyFromRows([][]any{{"yes"}}))
	panicMaybe(err)
}

func panicMaybe(err error) {
	if err != nil {
		panic(err)
	}
}
