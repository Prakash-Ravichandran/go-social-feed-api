package store

import (
	"context"
	"database/sql"
)

// app.store.posts.Create() / .update()
// app.store.users.Create() / .create()

type Storage struct {
	Posts interface {
		Create(context.Context) error
	}

	Users interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db: db},
		Users: &UserStore{db: db},
	}
}
