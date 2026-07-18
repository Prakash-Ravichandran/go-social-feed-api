package main

import (
	"log"

	db "github.com/Prakash-Ravichandran/go-social-feed-api/internal/db"
	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/env"
	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
	_ "github.com/lib/pq"
)

type dbconfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func main() {
	cfg := config{
		Addr: env.GetString("ADDR", ":8080"),
		db: dbconfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialfeed?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("DB connection pool established successfull")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
