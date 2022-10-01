package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	doStdlib()
	doNative()
}

func doStdlib() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "stdlib: Unable to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	var id string
	err = db.QueryRow("SELECT $1::text", 2).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "stdlib: QueryRow failed: %v\n", err)
		return
	}

	fmt.Println("stdlib:", id)
}

func doNative() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "native: Unable to connect to database: %v\n", err)
		return
	}
	defer conn.Close(context.Background())

	var id string
	err = conn.QueryRow(context.Background(), "SELECT $1::text", 2).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "native: QueryRow failed: %v\n", err)
		return
	}

	fmt.Println("native:", id)
}
