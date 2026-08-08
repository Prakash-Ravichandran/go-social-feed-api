package main

import "net/http"

type CreatePostPayload struct {
	Tilte   string `json:"title"`
	Content string `json:"content"`
}

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

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var postPayload CreatePostPayload

	if err := ReadJSON(w, http.StatusOK, r, &postPayload); err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteJSON(w, http.StatusCreated, postPayload)
}
