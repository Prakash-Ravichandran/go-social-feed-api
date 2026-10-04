package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

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

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		userIdInt64, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.Users.GetById(ctx, userIdInt64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromCtx(r) // user retrived from /users/{userId}

	// Revert back to auth userID from ctx
	var payload FollowUser
	// get and store user_id from the body { "user_id" : 199 } to &payload { user_id : 199}
	// payload.UserID is the autheticated user_id for now
	if err := ReadJSON(w, http.StatusOK, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	// who clicks follow, who is followed
	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unFollowerUser := getUserFromCtx(r)

	// TODO: Revert back to auth userID from ctx
	var payload FollowUser
	// get and store user_id from the endpoint users/{userId} to &payload { user_id : /1}
	if err := ReadJSON(w, http.StatusOK, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	// who clicks unfollow, who is unfollowed
	if err := app.store.Followers.UnFollow(ctx, unFollowerUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
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
