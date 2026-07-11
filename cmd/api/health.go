package main

import "net/http"

func (app *application) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all good"))
}
