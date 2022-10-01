package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	t1 := time.Now()
	_, err = conn.Exec(ctx, `select 1, pg_sleep(60)`)
	if err == nil {
		log.Fatal("expected error but did not receive")
	}
	t2 := time.Now()

	fmt.Println(err, t2.Sub(t1))
}
