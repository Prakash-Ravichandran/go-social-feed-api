package main

import (
	"net/http"
	"time"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	Addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recovers us from panic

	// handlers
	r.Get("/health", app.handleHealth)

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Println("servers has started running on", "addr", app.config.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	return nil
}
