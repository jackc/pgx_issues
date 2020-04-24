package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	var connCount = int(30)
	var queriesPerConn = int(1000)

	n, err := strconv.ParseInt(os.Getenv("CONN_COUNT"), 10, 32)
	if err == nil {
		connCount = int(n)
	}
	n, err = strconv.ParseInt(os.Getenv("QUERIES_PER_CONN"), 10, 32)
	if err == nil {
		queriesPerConn = int(n)
	}

	log.Printf("connCount: %d, queriesPerConn: %d\n", connCount, queriesPerConn)

	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	config.MaxConns = int32(connCount)
	config.ConnConfig.PreferSimpleProtocol = true

	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbpool.Close()

	doneChan := make(chan struct{})

	for i := 0; i < connCount; i++ {
		go func() {
			stressTest(dbpool, queriesPerConn)
			doneChan <- struct{}{}
		}()
	}

	for i := 0; i < connCount; i++ {
		<-doneChan
	}

	log.Println("No cancel errors detected.")
}

func stressTest(dbpool *pgxpool.Pool, queriesPerConn int) {
	for i := 0; i < queriesPerConn; i++ {
		n := rand.Float64()
		if n < 0.05 {
			// Slow query that is never canceled
			rows, _ := dbpool.Query(context.Background(), "select n, pg_sleep(0.001) from generate_series(1, 200) n")
			rows.Close()
			if rows.Err() != nil {
				log.Fatalf("A slow query that should never be canceled failed: %v", rows.Err())
			}
		} else if n < 0.9 {
			// Fast query that is never canceled
			rows, _ := dbpool.Query(context.Background(), "select n from generate_series(1, 100) n")
			rows.Close()
			if rows.Err() != nil {
				log.Fatalf("A fast query that should never be canceled failed: %v", rows.Err())
			}
		} else {
			// Query that is canceled
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
			rows, _ := dbpool.Query(ctx, "select pg_sleep(10)")
			rows.Close()
			if rows.Err() == nil {
				log.Fatalf("A query that should have been canceled was not")
			} else if !pgconn.Timeout(rows.Err()) {
				log.Fatalf("A query that should have been canceled did not get a timeout error")
			}
			cancel()
		}
	}
}
