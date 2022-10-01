package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/jackc/pgconn"
)

func main() {
	pgConn, err := pgconn.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("pgconn failed to connect:", err)
	}
	defer pgConn.Close(context.Background())

	for i := 0; i <= math.MaxUint8; i++ {
		result := pgConn.ExecParams(context.Background(), `SELECT $1::"char", $1::"char"`, [][]byte{{byte(i)}}, nil, []int16{1}, []int16{0, 1})
		for result.NextRow() {
			fmt.Printf("Text: %v, Binary: %v\n", result.Values()[0], result.Values()[1])
		}
		_, err = result.Close()
		if err != nil {
			log.Fatalln("failed reading result:", err)
		}

	}

}
