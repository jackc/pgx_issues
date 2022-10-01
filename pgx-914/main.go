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

	_, err = db.Exec("create temporary table t (id serial primary key, ip inet not null);")
	if err != nil {
		log.Fatal(err)
	}

	var paramIP pgtype.Inet
	err = paramIP.Set("10.0.0.0/16")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into t (ip) values ($1);", paramIP)
	if err != nil {
		log.Fatal(err)
	}

	var resultIP pgtype.Inet
	err = db.QueryRow("select ip from t").Scan(&resultIP)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultIP.IPNet)
}
