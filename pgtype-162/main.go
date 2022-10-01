package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type myTime struct {
	time.Time
}

// Scan implements the sql.Scanner interface
func (t *myTime) Scan(src interface{}) error {
	*t = myTime{Time: src.(time.Time)}
	return nil
}

func main() {
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	var t myTime
	err = conn.QueryRow(context.Background(), "SELECT NOW();").Scan(&t)
	if err != nil {
		panic(err)
	}

	var tPtr *myTime
	err = conn.QueryRow(context.Background(), "SELECT NOW();").Scan(&tPtr)
	if err != nil {
		panic(err)
	}
}
