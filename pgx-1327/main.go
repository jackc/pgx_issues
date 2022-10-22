package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	listener *pgx.Conn // dedicated notifications listener
	events   chan string
	*pgxpool.Pool
}

func PreparePool(ctx context.Context, connectionString string, listenChannels []string) (*Pool, error) {
	var p Pool
	p.events = make(chan string, 20)
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	listenerConfig := poolConfig.ConnConfig.Copy()
	listenerConfig.OnNotification = func(_ *pgconn.PgConn, n *pgconn.Notification) {
		p.events <- fmt.Sprintf("notification: %s, payload: %s", n.Channel, n.Payload)
	}

	listenerConfig.AfterConnect = func(ctx context.Context, conn *pgconn.PgConn) error {
		go func() {
			if len(listenChannels) == 0 {
				return
			}
			var q string
			for _, ch := range listenChannels {
				q = fmt.Sprintf("%slisten %s; ", q, ch)
			}
			res := conn.Exec(ctx, q)
			res.Close()

			for {
				err := conn.WaitForNotification(ctx)
				if err != nil {
					p.events <- err.Error() + "\nnotifications listener stopped!"
					break
				}
			}

		}()
		return nil
	}

	p.listener, err = pgx.ConnectConfig(ctx, listenerConfig)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.OnNotice = func(_ *pgconn.PgConn, n *pgconn.Notice) {
		p.events <- (*pgconn.PgError)(n).Error()
	}

	p.Pool, err = pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		p.listener.Close(ctx)
		return nil, err
	}

	return &p, nil
}

func tryConnectionString(id int, cs string) {
	fmt.Println("--- ", id)
	ctx, cancel := context.WithCancel(context.Background())

	p, err := PreparePool(ctx, cs, []string{"test"})
	if err != nil {
		fmt.Println(err)
		cancel()
		return
	}

	go func() {
		for {
			select {
			case e := <-p.events:
				fmt.Println(e)
			case <-ctx.Done():
				fmt.Println("finished")
				return
			}
		}
	}()

	p.AcquireFunc(ctx, func(c *pgxpool.Conn) error {
		_, err := c.Exec(ctx, `do
		$$
		begin
			raise notice 'raise notice test';
			perform pg_notify('test', 'payload');
			raise warning 'raise warning test';
		end
		$$`)
		return err
	})

	time.Sleep(100 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func main() {
	cs := ""
	tryConnectionString(1, cs)
	tryConnectionString(2, cs+" connect_timeout=60")
	tryConnectionString(3, cs+"\t\n")
}
