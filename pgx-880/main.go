package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select $1", pgx.QuerySimpleProtocol(true), "Hello, world")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
