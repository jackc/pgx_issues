package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connConfig, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db := stdlib.OpenDB(*connConfig, stdlib.OptionAfterConnect(
		func(ctx context.Context, conn *pgx.Conn) error {
			info := conn.TypeMap()
			pgxUUID.Register(info)
			return nil
		}),
	)

	var u uuid.UUID
	u2 := uuid.New()
	err = db.QueryRow(
		`SELECT $1::uuid`,
		u2,
	).Scan(&u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
}
