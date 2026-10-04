package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// app.store.posts.Create() / .update()
// app.store.users.Create() / .create()
// app.store.Comments.GetPostById()

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exist")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		UpdateById(context.Context, *Post) (*Post, error)
		DeleteById(context.Context, int64) error
	}

	Users interface {
		Create(context.Context, *User) error
		GetById(context.Context, int64) (*User, error)
		DeleteById(context.Context, int64) error
	}

	Comments interface {
		GetPostById(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}

	// in real world Follow is not real entity, due to separation of concerns keeping outside User Model
	Followers interface {
		Follow(ctx context.Context, followerID int64, userID int64) error
		UnFollow(ctx context.Context, unFollowerID int64, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db: db},
		Users:     &UserStore{db: db},
		Comments:  &CommentStore{db: db},
		Followers: &FollowerStore{db: db},
	}
}
