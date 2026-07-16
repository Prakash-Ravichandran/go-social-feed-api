package main

import (
	"log"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/env"
	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
)

func main() {

	cfg := config{
		Addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
