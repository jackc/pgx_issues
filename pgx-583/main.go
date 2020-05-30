package main

import (
	"context"
	"database/sql/driver"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type colorType int

const (
	blue colorType = iota
)

func (t colorType) Value() (driver.Value, error) {
	return "blue", nil
}

var ErrInvalid = errors.New("invalid type")

func (t *colorType) Scan(src interface{}) error {
	strVal, ok := src.(string)
	if !ok || strVal != "blue" {
		return ErrInvalid
	}
	*t = blue
	return nil
}

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())

	var color colorType
	err = conn.QueryRow(context.Background(), "select single_color from blah where id=1").Scan(&color)
	if err != nil {
		log.Fatalln(err)
	}

}
