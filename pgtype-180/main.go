package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var ip pgtype.Inet

	err = db.QueryRow("select '127.0.0.1'::inet").Scan(&ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ip)
}
