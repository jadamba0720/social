package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		GetById(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}
	User interface {
		GetById(context.Context, int64) (*User, error)
		Create(context.Context, *User) error
	}
	Comment interface {
		GetByPostID(context.Context, int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}
	Follow interface {
		Follow(ctx context.Context, followId, userID int64) error
		UnFollow(ctx context.Context, followId, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:   &PostStore{db},
		User:    &UserStore{db},
		Comment: &CommentStore{db},
		Follow:  &FollowerStore{db},
	}
}
