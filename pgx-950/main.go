package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	scanArgs := []interface{}{}
	sql := &strings.Builder{}
	sql.WriteString("select ")
	for i := 0; i < 1660; i++ {
		if i != 0 {
			sql.WriteString(", ")
		}
		fmt.Fprintf(sql, "1 as column_with_really_really_really_really_long_name_%d", i)
		var n int32
		scanArgs = append(scanArgs, &n)
	}

	sql.WriteString(", pg_sleep(5)")
	scanArgs = append(scanArgs, nil)

	log.Println("Call QueryRow")
	row := conn.QueryRow(context.Background(), sql.String())
	log.Println("QueryRow returned")

	log.Println("Call Scan")
	err = row.Scan(scanArgs...)
	log.Println("Scan returned")
	if err != nil {
		log.Fatalln(err)
	}
}
