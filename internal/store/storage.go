package store

import (
	"context"
	"database/sql"
)

// app.store.posts.Create() / .update()
// app.store.users.Create() / .create()
// app.store.Comments.GetPostById()

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		UpdateById(context.Context, *Post) (*Post, error)
		DeleteById(context.Context, int64) error
	}

	Users interface {
		Create(context.Context, *User) error
	}

	Comments interface {
		GetPostById(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db: db},
		Users:    &UserStore{db: db},
		Comments: &CommentStore{db: db},
	}
}
