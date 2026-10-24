package main

import (
	"net/http"

	"github.com/Prakash-Ravichandran/go-social-feed-api/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	fq := store.PaginatedFeedQuery{
		Limit:  20, // default value 20
		Offset: 0,
		Sort:   "desc", // default desc sort
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(201), fq)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) dummyFeedHandler(w http.ResponseWriter, r *http.Request) {
	app.jsonResponse(w, http.StatusOK, "All Good")
}
