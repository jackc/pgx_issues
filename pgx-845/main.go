package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"time"

	// profiler here
	"net/http"
	_ "net/http/pprof"

	"github.com/jackc/pgx/v4/pgxpool"
)

var content string

func main() {
	contentBytes := make([]byte, 50*1024*1024)
	_, err := rand.Read(contentBytes)
	if err != nil {
		log.Fatalln(err)
	}
	content = hex.EncodeToString(contentBytes)

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	start := time.Now()
	executeLargeQuery()
	// runtime.GC()
	log.Printf("executeLargeQuery completed in %v\n", time.Since(start))
	log.Println(http.ListenAndServe(":12345", nil))
}

func executeLargeQuery() {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer pool.Close()
	for i := 0; true; i++ {
		rows, err := pool.Query(ctx, "select $1, $2::int", content, i)
		if err != nil {
			log.Fatalln(err)
		}
		rows.Close()
		if rows.Err() != nil {
			log.Fatalln(rows.Err())

		}
		log.Printf("iteration: %v\n", i)
	}
}
