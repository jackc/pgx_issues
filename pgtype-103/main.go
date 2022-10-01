package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type Optional struct {
	a int
	b int
}

type Complex struct {
	n   int
	opt *Optional
}

// These functions make Optional implement pgtype.Value, though they are not used in this example.
func (dst *Optional) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to Optional", src)
}
func (dst *Optional) Get() interface{} {
	return dst
}
func (src *Optional) AssignTo(dst interface{}) error {
	return fmt.Errorf("cannot assign %v to %T", src, dst)
}

// EncodeBinary encodes an Optional into binary.
func (src *Optional) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
	if src == nil {
		return nil, nil
	}

	b, err := (&pgtype.CompositeFields{
		&pgtype.Int8{Int: int64(src.a), Status: pgtype.Present},
		&pgtype.Int8{Int: int64(src.b), Status: pgtype.Present},
	}).EncodeBinary(ci, buf)
	if err != nil {
		return nil, fmt.Errorf("could not encode Optional: %w", err)
	}
	return b, nil
}

func (dst *Optional) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		return errors.New("NULL values can't be decoded. Scan into a &*Optional to handle NULLs")
	}

	err := (pgtype.CompositeFields{&dst.a, &dst.b}).DecodeBinary(ci, src)
	if err != nil {
		return fmt.Errorf("could not decode composite fields: %w", err)
	}
	return nil
}

func main() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Register Optional's OID. Otherwise we cannot scan it.
	var optionalOID uint32
	err = conn.QueryRow(context.Background(), `select 'optional'::regtype::oid`).Scan(&optionalOID)
	if err != nil {
		log.Fatal(err)
	}
	// conn.ConnInfo().RegisterDataType(pgtype.DataType{Value: &Optional{}, Name: "optional", OID: optionalOID})

	compNil := Complex{n: 4, opt: nil}
	_, err = conn.Exec(context.Background(), "INSERT INTO complex(n, opt) VALUES ($1, $2)", compNil.n, compNil.opt)
	if err != nil {
		log.Fatal(err)
	}

	var oNil *Optional
	err = conn.QueryRow(context.Background(), "SELECT opt FROM complex WHERE n=$1 LIMIT 1", compNil.n).Scan(&oNil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("oNil:", oNil)

	comp := Complex{n: 3, opt: &Optional{a: 1, b: 2}}
	_, err = conn.Exec(context.Background(), "INSERT INTO complex(n, opt) VALUES ($1, $2)", comp.n, comp.opt)
	if err != nil {
		log.Fatal(err)
	}

	var o *Optional
	err = conn.QueryRow(context.Background(), "SELECT opt FROM complex WHERE n=$1 LIMIT 1", pgx.QueryResultFormats{pgx.BinaryFormatCode}, comp.n).Scan(&o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("o:", o)
}
