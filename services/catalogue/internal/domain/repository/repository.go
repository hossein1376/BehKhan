package repository

import (
	"context"
	"database/sql"
)

type QueryFunc = func(r *Repo) error

type Querier interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type Pool interface {
	Query(ctx context.Context, f QueryFunc) error
	Close() error
}

type Tables struct {
	Books   BookRepository
	Authors AuthorRepository
}

type Repo struct {
	Tables
	Querier
}
