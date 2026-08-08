package main

import (
	"net/http"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
)

type CreatePostPayload struct {
	Tilte   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
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

	// capture user payload info into postpayload
	if err := ReadJSON(w, http.StatusOK, r, &postPayload); err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()

	posts := &store.Post{
		Title:   postPayload.Tilte,
		Content: postPayload.Content,
		Tags:    postPayload.Tags,
		// TODO: change after auth
		UserID: "1",
	}

	if err := app.store.Posts.Create(ctx, posts); err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := WriteJSON(w, http.StatusCreated, posts); err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
