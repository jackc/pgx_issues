package main

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connConfig, _ := pgx.ParseConfig("")
	connConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	conn, err := sql.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	//conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(`drop table if exists pgx514;`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(`create table pgx514 (id serial primary key, data jsonb not null);`)
	if err != nil {
		log.Fatal(err)
	}

	dataJSON := &pgtype.Text{String: `{"foo": "bar"}`, Valid: true}
	commandTag, err := conn.Exec("insert into pgx514(data) values($1)", dataJSON)
	if err == nil {
		log.Println("pgtype.JSON", commandTag)
	} else {
		log.Println("pgtype.JSON", err)
	}

	dataBytes := []byte(`{"foo": "bar"}`)
	commandTag, err = conn.Exec("insert into pgx514(data) values($1)", dataBytes)
	if err == nil {
		log.Println("[]byte", commandTag)
	} else {
		log.Println("[]byte", err)
	}
}
