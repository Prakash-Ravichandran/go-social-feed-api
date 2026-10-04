package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// Model
type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type FollowerStore struct {
	db *sql.DB
}

func (fs *FollowerStore) Follow(ctx context.Context, followerID int64, userID int64) error {
	query := `
    INSERT INTO followers 
	(user_id, follower_id) 
	VALUES ($1, $2)
   `
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := fs.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		// TODO: need to revisit to return a proper error
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}

	return nil
}

func (fs *FollowerStore) UnFollow(ctx context.Context, unFollowerID int64, userID int64) error {
	query := `
   DELETE FROM followers 
   where user_id = $1 AND follower_id = $2  
   `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := fs.db.ExecContext(ctx, query, userID, unFollowerID)
	if err != nil {
		// TODO: need to revisit to return a proper error
		return err
	}
	return nil
}
