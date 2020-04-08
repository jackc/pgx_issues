package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"golang.org/x/sync/errgroup"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		log.Fatalln(err)
	}

	g, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 50; i++ {
		g.Go(func() error {
			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				log.Println(err)
				return err
			}

			rows, err := tx.QueryContext(ctx, "SELECT table_name FROM information_schema.tables")
			if err != nil {
				log.Println(err)
				return err
			}

			if !rows.Next() {
				panic("no rows")
			}
			rows.Close()
			if rows.Err() != nil {
				log.Println("rows.Err()", err)
			}

			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}

			return nil
		})
	}

	_ = g.Wait()
}
