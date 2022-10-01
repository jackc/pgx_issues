package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	insert := func(v [768]float64) (*int, error) {
		insertVectorSQL := `
		INSERT INTO vectors
			(vector)
		VALUES
			($1)
		RETURNING
			_id
		`
		var vecId int
		err := conn.QueryRow(insertVectorSQL, pq.Array(v)).Scan(&vecId)
		if err != nil {
			return nil, err
		}
		return &vecId, err
	}

	a := [768]float64{}
	a[0] = 1
	i, err := insert(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("vector id: %d", i)

	b := [768]float64{}
	b[0] = 0
	i, err = insert(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("vector id: %d", i)
}
