package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	_, err = db.Exec("drop table if exists pgx748")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec("create table pgx748 (a interval not null)")
	if err != nil {
		log.Fatalln(err)
	}

	in := pgtype.Interval{}
	err = in.Set(time.Second * 5)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec("insert into pgx748 values($1)", &in)
	if err != nil {
		log.Fatalln(err)
	}

	var out pgtype.Interval
	err = db.QueryRow("select a from pgx748").Scan(&out)
	if err != nil {
		log.Fatalln(err)
	}

	var duration time.Duration
	err = out.AssignTo(&duration)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(duration)
}
