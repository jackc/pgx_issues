package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
)

func main() {
	pgConn, err := pgconn.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("pgconn failed to connect:", err)
	}
	defer pgConn.Close(context.Background())

	results, err := pgConn.Exec(context.Background(), "begin; SELECT 1; commit; create database test1;").ReadAll()
	if err != nil {
		log.Fatalln("pgconn failed to connect:", err)
	}

	fmt.Println(results)
}
