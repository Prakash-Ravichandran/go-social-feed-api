package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// Post - Model
type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    string   `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (ps *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	err := ps.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(
		&post.ID, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// app.store.Posts.GetById
func (ps *PostStore) GetById(ctx context.Context, id int64) (*Post, error) {
	query := `
    SELECT * FROM posts WHERE id = $1  
   `
	var post Post
	err := ps.db.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.Title, &post.UserID, &post.Content, &post.CreatedAt, pq.Array(&post.Tags), &post.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &Post{
		ID:        post.ID,
		Content:   post.Content,
		Title:     post.Title,
		UserID:    post.UserID,
		Tags:      post.Tags,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}
