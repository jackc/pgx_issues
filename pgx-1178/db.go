package main

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
)

type _DB struct {
	User   string
	Pass   string
	Addr   string
	Name   string
	Conn   *pgx.Conn
	Status bool
	AppCfg _Cfg
}

func (db *_DB) Connected() bool {
	return db.Status
}

func (db *_DB) Close() {
	db.Conn.Close(context.Background())
}

func (db *_DB) Connect() error {

	var err error
	db.Conn, err = pgx.Connect(context.Background(), "")

	if err != nil {
		db.Status = false
		return errors.New(err.Error())
	}

	db.Status = true

	return nil
}

func (db *_DB) LoadCfg(app *_App, print bool) error {

	err := db.Conn.QueryRow(context.Background(), "SELECT cfg::json FROM services WHERE name = $1", app.Name).Scan(&db.AppCfg)
	if err != nil {
		if print {
			log.Println(err)
		}
		return errors.New(err.Error())
	}

	log.Printf("dbCFG - %v", db.AppCfg)

	return nil
}
