package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	connStr := "user=postgres password=postgres dbname=postgres host=127.0.0.1 sslmode=disable"
	driver := "pgx"
	db, err := sql.Open(driver, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	fmt.Printf("ping err %v\n", err)
}
