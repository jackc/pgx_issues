package main

import (
	"context"
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

type customDate struct {
	t time.Time
}

func (d customDate) Value() (driver.Value, error) {
	return d.t.Format("2006-01-02"), nil
}

func (d *customDate) Scan(src interface{}) (err error) {
	if src == nil {
		d.t = time.Time{}
		return nil
	}

	switch v := src.(type) {
	case int64:
		d.t = time.Unix(v, 0).UTC()
	case float64:
		d.t = time.Unix(int64(v), 0).UTC()
	case string:
		d.t, err = time.Parse("2006-01-02", v)
	case []byte:
		d.t, err = time.Parse("2006-01-02", string(v))
	case time.Time:
		d.t = v
	default:
		err = fmt.Errorf("failed to scan type '%T' into date", src)
	}
	return err
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	dateIn := []customDate{{t: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)}}

	var dateOut []customDate
	err = conn.QueryRow(ctx, "select $1::date[]", dateIn).Scan(&dateOut)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dateOut[0].t.GoString())
}
