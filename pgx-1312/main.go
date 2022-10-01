package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	var arr pgtype.Array[string]
	m := pgtype.NewMap()
	err = db.QueryRow(
		`SELECT NULL::text[]`,
	).Scan(m.SQLScanner(&arr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)
}
