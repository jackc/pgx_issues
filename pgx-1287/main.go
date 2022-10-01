package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	os.Setenv("PGSERVICE", "training-staging")
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	closeCtx, _ := context.WithTimeout(context.Background(), 300*time.Second)

	// Your code here...
	for i := 0; closeCtx.Err() == nil; i++ {

		qctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		if errPing := conn.Ping(qctx); errPing == nil {
			fmt.Println(i, "DB is here")
		} else {
			fmt.Println(i, "DB is not here")
		}
		// time.Sleep(time.Second)
	}

	fmt.Println("leaving")
}
