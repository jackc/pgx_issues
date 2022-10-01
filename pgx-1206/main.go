package main

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func registerHstore(conn *pgx.Conn) {
	var hstoreOID uint32
	err := conn.QueryRow(context.Background(), "select t.oid from pg_type t where t.typname='hstore';").Scan(&hstoreOID)
	if err != nil {
		panic(err)
	}
	conn.ConnInfo().RegisterDataType(pgtype.DataType{Value: &pgtype.Hstore{}, Name: "hstore", OID: hstoreOID})
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	registerHstore(conn)

	conn2, err := pgx.Connect(ctx, "")
	if err != nil {
		panic(err)
	}
	defer conn2.Close(ctx)
	registerHstore(conn2)

	rows, err := conn2.Query(ctx, "select * from test2")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	_, err = conn.CopyFrom(
		ctx,
		[]string{"test"},
		[]string{"hs", "hs2"},
		rows,
	)
	if err != nil {
		panic(err)
	}
}
