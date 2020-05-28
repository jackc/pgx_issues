package main

import (
	"database/sql"

	"github.com/jackc/pgx/v4/stdlib"
)

func useAcquireConn(db *sql.DB) error {
	conn, err := stdlib.AcquireConn(db)
	if err != nil {
		return err
	}

	return stdlib.ReleaseConn(db, conn)
}
