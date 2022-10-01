package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	ip := net.ParseIP("1.1.1.1")
	fmt.Println("Before", ip, ip.String(), len(ip))

	var str string
	err = db.QueryRow(
		`SELECT $1::inet`,
		ip,
	).Scan(&str)
	if err != nil {
		log.Fatal(err)
	}

	ip2 := net.ParseIP(str)
	fmt.Println("str", str)
	fmt.Println("After", ip2, ip2.String(), len(ip2))
}
