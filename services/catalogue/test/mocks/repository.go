package mocks

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repository"
	"github.com/hossein1376/BehKhan/catalogue/pkg/errs"
)

type Pool struct{}

func (Pool) Query(_ context.Context, f repository.QueryFunc) error {
	repo := &repository.Repo{
		Tables: repository.Tables{
			Books:   BookRepository{},
			Authors: AuthorRepository{},
		},
		Querier: Querier{},
	}
	return f(repo)
}

func (Pool) Close() error {
	return nil
}

type Querier struct{}

func (Querier) ExecContext(_ context.Context, _ string, _ ...any) (sql.Result, error) {
	return nil, nil
}

func (Querier) QueryContext(_ context.Context, _ string, _ ...any) (*sql.Rows, error) {
	return nil, nil
}

type BookRepository struct{}

func (BookRepository) Create(context.Context, entity.Book) error {
	return nil
}

func (BookRepository) GetByID(_ context.Context, id entity.BookID) (*entity.Book, error) {
	if id < 1 {
		return nil, errs.NotFound(fmt.Errorf("not found"))
	}
	return &entity.Book{
		ID:    id,
		Title: strconv.Itoa(int(id)),
	}, nil
}

type AuthorRepository struct{}
