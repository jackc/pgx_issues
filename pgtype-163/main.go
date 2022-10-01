package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var j pgtype.JSON

	err = db.QueryRow(`select null::json`).Scan(&j)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(j.Status == pgtype.Null)
}
