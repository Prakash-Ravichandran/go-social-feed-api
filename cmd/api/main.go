package main

import (
	"log"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/env"
)

func main() {

	cfg := config{
		Addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
