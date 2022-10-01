package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgconn"
)

func main() {
	pgConn, err := pgconn.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("pgconn failed to connect:", err)
	}
	defer pgConn.Close(context.Background())

	result := pgConn.ExecParams(context.Background(), "SELECT email FROM users WHERE id=$1", [][]byte{[]byte("123")}, nil, nil, nil)
	for result.NextRow() {
		fmt.Println("User 123 has email:", string(result.Values()[0]))
	}
	_, err = result.Close()
	if err != nil {
		log.Fatalln("failed reading result:", err)
	}
}
