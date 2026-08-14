package main

import (
	"net/http"

	"log"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	WriteErrorJSON(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	WriteErrorJSON(w, http.StatusBadRequest, err.Error())
}
