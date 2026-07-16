package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) Create(ctx context.Context) error {
	return nil
}
