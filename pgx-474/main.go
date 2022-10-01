package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var n int
	err = db.QueryRow("select 1/0").Scan(&n)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.DivisionByZero:
				log.Println("correctly matched division by zero error")
				os.Exit(0)
			}
		}
	}
	log.Fatalln("did not match error", err)
}
