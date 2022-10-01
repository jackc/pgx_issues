package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())

	conn.Exec(context.Background(), `
	drop table if exists transactions cascade;
	create table if not exists transactions (
			id bigserial primary key,
			constraint id_gte_zero check (id >= 0),

			source_amount decimal not null,
			constraint source_amount_gt_zero
					check (source_amount > 0)
	);

	insert into transactions (source_amount) values (10);
	`)

	rows, err := conn.Query(context.Background(), "SELECT id, source_amount from transactions")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatalln(err)
		}
		num := values[1].(pgtype.Numeric)
		var f float64
		err = num.AssignTo(&f)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Int: %v -- Exp: %v -- Float: %v\n", num.Int, num.Exp, f)
	}

	if rows.Err() != nil {
		log.Fatalln(err)
	}
}
