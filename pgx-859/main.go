package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"unsafe"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	var a, b, c []byte
	var pool *pgxpool.Pool
	var conn *pgxpool.Conn

	pool, _ = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn, _ = pool.Acquire(context.Background())

	a = []byte{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47}
	b = []byte{64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95}
	c = []byte{112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143}

	err := conn.QueryRow(context.Background(), "select '\\x404142434445464748494A4B4C4D4E4F505152535455565758595A5B5C5D5E5F'::bytea, '\\x707172737475767778797A7B7C7D7E7F808182838485868788898A8B8C8D8E8F'::bytea").Scan(&b, &c)
	if err != nil {
		log.Fatal(err)
	}

	// b = b[:len(b):len(b)]

	fmt.Printf("%v\n%v\n%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)), (*reflect.SliceHeader)(unsafe.Pointer(&b)), (*reflect.SliceHeader)(unsafe.Pointer(&c)))

	fmt.Println("c1", c)

	_ = append(b, a...)

	fmt.Println("c2", c)

	conn.Release()
	pool.Close()

}
