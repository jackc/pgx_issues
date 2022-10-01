package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	poolConf, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	poolConf.BeforeAcquire = func(ctx context.Context, conn *pgx.Conn) bool {
		log.Printf("BeforeAcquire conn: %p\n", conn)
		return true
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConf)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()

	log.Printf("Acquire conn: %p\n", conn.Conn())
}
