package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	config.PreferSimpleProtocol = true

	busyConn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer busyConn.Close(context.Background())

	protocolPID, sqlPID, secretKey, err := getConnInfo(busyConn)
	if err != nil {
		log.Fatalf("failed to get busyConn info: %v\n", err)
	}
	log.Printf("[busyConn]: protocolPID=%d, sqlPID=%d, secretKey=%d\n", protocolPID, sqlPID, secretKey)

	cancelConn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalln(err)
	}
	defer cancelConn.Close(context.Background())

	protocolPID, sqlPID, secretKey, err = getConnInfo(cancelConn)
	if err != nil {
		log.Fatalf("failed to get cancelConn info: %v\n", err)
	}
	log.Printf("[cancelConn]: protocolPID=%d, sqlPID=%d, secretKey=%d\n", protocolPID, sqlPID, secretKey)

	doneChan := make(chan struct{})

	go func() {
		sql := `select pg_sleep(10)`
		log.Printf("[busyConn] executing: %s\n", sql)
		_, err := busyConn.Exec(context.Background(), sql)
		if err == nil {
			log.Printf("[busyConn] was not interrupted\n")
		} else {
			log.Printf("[busyConn] err: %v\n", err)
		}

		doneChan <- struct{}{}
	}()

	go func() {
		log.Printf("[cancelConn]: waiting for other conn to run query...\n")
		time.Sleep(2 * time.Second)
		log.Printf("[cancelConn]: sending CancelRequest\n")
		err := cancelConn.PgConn().CancelRequest(context.Background())
		log.Printf("[cancelConn]: CancelRequest err: %v\n", err)
		doneChan <- struct{}{}
	}()

	<-doneChan
	<-doneChan
}

func getConnInfo(conn *pgx.Conn) (protocolPID uint32, sqlPID uint32, secretKey uint32, err error) {
	err = conn.QueryRow(context.Background(), "select pg_backend_pid()").Scan(&sqlPID)
	return conn.PgConn().PID(), sqlPID, conn.PgConn().SecretKey(), err
}
