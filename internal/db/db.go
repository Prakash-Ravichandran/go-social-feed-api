package db

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func New(addr string, maxOpenConns int, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		log.Panic(err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(duration)

	// if conection takes more than 5 seconds, then we will have a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// pingcontext is the one that establishes the connection
	if err = db.PingContext(ctx); err != nil {
		log.Println("returning from pingcontext")
		return nil, err
	}

	return db, nil
}
