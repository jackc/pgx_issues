package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var n int64
	err = db.QueryRow("select cardinality($1::text[])", []string{"a", "b", "c"}).Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(n)
}
