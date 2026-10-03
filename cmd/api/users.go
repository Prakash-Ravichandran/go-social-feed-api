package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) usersHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "Users - All Good",
		"env":     app.config.env,
		"version": version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.badRequestResponse(w, r, err)
	}
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
	}

	ctx := r.Context()

	user, err := app.store.Users.GetById(ctx, userIdInt64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	app.jsonResponse(w, http.StatusOK, user)
}

// Revisit to for the below error
/**
2026/10/03 12:27:42 servers has started running on addr :3000
2026/10/03 12:27:48 ERROR inside DeleteById for userId 90: pq: update or delete on table "users" violates foreign key constraint "fk_user" on table "posts" (23503)
2026/10/03 12:27:48 internal error method DELETE path /users/90 error pq: update or delete on table "users" violates foreign key constraint "fk_user" on table "posts" (23503)
2026/10/03 12:27:48 [Prakash/CGLJqFZZ2F-000001] "DELETE http://127.0.0.1:3000/users/90 HTTP/1.1" from 127.0.0.1 - 500 45B in 5.001ms
**/
func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Users.DeleteById(ctx, userIdInt64); err != nil {
		log.Printf("ERROR inside DeleteById for userId %d: %v", userIdInt64, err)
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
