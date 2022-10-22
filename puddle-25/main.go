package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := pgxpool.New(ctx, "")
	if err != nil {
		log.Panicln(err)
	}
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(8)

	loop := func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			g.Go(func() error {
				log.Println("Going to db.Exec")
				_, err := db.Exec(ctx, "select pg_sleep(0.2);")
				return err
			})
		}
	}

	loop()
	db.Close()
}
