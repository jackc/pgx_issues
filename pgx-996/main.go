package main

import (
	"context"
	"errors"
	"fmt"
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

	var pid uint32
	err = conn.QueryRow(context.Background(), "select pid from pg_stat_activity where false").Scan(&pid)
	fmt.Printf("equal? %v\n", err.Error() == pgx.ErrNoRows.Error())
	fmt.Printf("equal? %v\n", errors.Is(err, pgx.ErrNoRows))
}
