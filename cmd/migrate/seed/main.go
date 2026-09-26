package main

import (
	"log"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/db"
	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/env"
	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
)

func main() {

	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialfeed?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal("Error Creating new db", err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
