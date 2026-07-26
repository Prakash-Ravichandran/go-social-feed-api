package main

import (
	"net/http"
)

func (app *application) handleHealth(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "All Good With Health",
		"env":     app.config.env,
		"version": version,
	}

	if err := WriteJSON(w, http.StatusOK, data); err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, "err.Error()")
	}
}
