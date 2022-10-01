package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

// ----- SETTINGS -----

// THe number of rows to insert
const INSERT_COUNT = 1000 // << Less 319 causes err?

func main() {

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Prepare the database to insert the test datter
	var setup pgx.Batch
	setup.Queue(sqlCreateTable)

	var res = conn.SendBatch(context.Background(), &setup)
	_, err = res.Exec()
	if err != nil {
		panic(err)
	}

	err = res.Close()
	if err != nil {
		panic(err)
	}

	// Prepare the test data which is three columns
	// of random strings
	log.Println("Generating Test Data:", INSERT_COUNT)
	var dataToInsert [][]interface{}

	for i := 0; i < INSERT_COUNT; i++ {
		var row []interface{}

		row = append(row, RandomString(64))
		row = append(row, RandomString(64))
		row = append(row, RandomString(64))

		dataToInsert = append(dataToInsert, row)
	}

	// Do the copy
	log.Println("Begin Insert Rows:", len(dataToInsert))
	var ident = pgx.Identifier{"three_strings"}
	var cols = []string{"col_1", "col_2", "col_3"}
	numRowsCopied, err := conn.CopyFrom(context.Background(), ident, cols, pgx.CopyFromRows(dataToInsert))
	if err != nil {
		log.Fatal(err)
	}

	// Ask Postgres how many rows are in the table
	var dbRowCount int
	err = conn.QueryRow(context.Background(), sqlCountRows).Scan(&dbRowCount)
	if err != nil {
		panic(err)
	}

	log.Println("----- RESULTS -----")
	log.Println("Should have written:", len(dataToInsert))
	log.Println("PGX says we wrote: ", numRowsCopied)
	log.Println("Actual rows in db: ", dbRowCount)
	log.Println("--------------------")

}

// Implements the CopyFrom interface
type Copier struct {
	data   [][]string
	err    error
	idx    int
	values []interface{}
}

// Moves the cursor and queues the next row
func (c *Copier) Next() bool {

	if c.idx >= len(c.data) {
		c.err = io.EOF
		return false
	}

	for _, cValue := range c.data[c.idx] {
		c.values = append(c.values, cValue)
	}

	c.idx = c.idx + 1
	return true
}

func (c *Copier) Err() error {
	return c.err
}

func (c *Copier) Values() ([]interface{}, error) {
	var vals = c.values
	c.values = nil
	return vals, c.err
}

// ----- Thanks Google
// ----- https://www.calhoun.io/creating-random-strings-in-go/
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(charset))]
	}
	return string(b)
}

// ----- SQL Setup
const sqlCreateTable = `create temporary table three_strings
(
	col_1 text,
	col_2 text,
	col_3 text
);`

const sqlCountRows = `SELECT COUNT(*) from three_strings;`
