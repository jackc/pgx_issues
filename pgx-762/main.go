package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var n int
	err = db.QueryRow("select 1").Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	if n != 1 {
		log.Fatalln("Unexpected query result")
	}

	err = useAcquireConn(db)
	if err != nil {
		log.Fatalln(err)
	}
}
