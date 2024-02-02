package db

import (
	"context"
	"database/sql"
	"log"
)

var stringFormat = "2006-01-02 15:04:05"

func acquireConn(ctx context.Context, db *sql.DB) (*sql.Conn, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Println("Error creating connection!")
		log.Println(err.Error())
		return nil, err
	}
	return conn, nil
}
