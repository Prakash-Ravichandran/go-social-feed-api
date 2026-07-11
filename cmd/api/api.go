package main

import (
	"net/http"
	"time"

	"log"
)

type application struct {
	config config
}

type config struct {
	Addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", app.handleHealth)
	return mux
}

func (app *application) run(mux *http.ServeMux) error {

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
