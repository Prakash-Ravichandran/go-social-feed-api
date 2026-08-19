package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

type UpdatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) postsHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "Posts - All Good",
		"env":     app.config.env,
		"version": version,
	}

	if err := WriteJSON(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var postPayload CreatePostPayload

	// capture user payload info into postpayload
	if err := ReadJSON(w, http.StatusOK, r, &postPayload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// validate the field - required and max characters for title & content
	if err := Validate.Struct(postPayload); err != nil {
		ValidationErrors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("validation error: %s", ValidationErrors), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	posts := &store.Post{
		Title:   postPayload.Title,
		Content: postPayload.Content,
		Tags:    postPayload.Tags,
		// TODO: change after auth
		UserID: "1",
	}

	if err := app.store.Posts.Create(ctx, posts); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := WriteJSON(w, http.StatusCreated, posts); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostsById(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")

	postInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()
	// app.store.Posts.GetById -> PostStore implements the interface GetById
	posts, err := app.store.Posts.GetById(ctx, postInt64)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	WriteJSON(w, http.StatusOK, posts)
}

func (app *application) updatePostsById(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")

	postInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return
	}

	ctx := r.Context()

	var tempUpdatePost UpdatePostPayload

	// capture JSON from user and store to appropriate typed data structure
	if err := ReadJSON(w, http.StatusOK, r, &tempUpdatePost); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// validate the field - required and max characters for title & content
	if err := Validate.Struct(tempUpdatePost); err != nil {
		ValidationErrors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("validation error: %s", ValidationErrors), http.StatusBadRequest)
		return
	}

	tempPost := &store.Post{
		Title:   tempUpdatePost.Title,
		Content: tempUpdatePost.Content,
		Tags:    tempUpdatePost.Tags,
	}

	post, err := app.store.Posts.UpdateById(ctx, postInt64, tempPost)
	if err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, post)
}

func (app *application) deletePostsById(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")

	postInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	if err := app.store.Posts.DeleteById(ctx, postInt64); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			WriteErrorJSON(w, http.StatusNotFound, "post not found")
		default:
			WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// https://stackoverflow.com/questions/2342579/http-status-code-for-update-and-delete
	// a 204 response - 204 (No Content) if the action has been enacted but the response does not include an entity.
	w.WriteHeader(http.StatusNoContent)
}
