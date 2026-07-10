package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Custom errors enum
var (
	ErrNotFound = errors.New("resource not found") // 404
	ErrConflict = errors.New("conflict error")     // 409
)

var QueryTimeoutDuration = time.Second * 5 // timeout duration

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		Update(context.Context, *Post) error
		Delete(context.Context, int64) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
