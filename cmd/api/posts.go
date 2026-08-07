package main

import "net/http"

func (app *application) postsHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "Posts - All Good",
		"env":     app.config.env,
		"version": version,
	}

	if err := WriteJSON(w, http.StatusOK, data); err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, "err.Error()")
	}
}
