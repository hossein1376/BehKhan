package booksrvc

import (
	"context"
	"fmt"

	"github.com/hossein1376/BehKhan/catalogue/internal/application/mapper"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repository"
	"github.com/hossein1376/BehKhan/catalogue/pkg/errs"
)

type BookSrvc struct {
	db repository.Pool
}

func NewBookSrvc(db repository.Pool) BookSrvc {
	return BookSrvc{db: db}
}

func (c BookSrvc) Create(ctx context.Context, title string) error {
	book, err := mapper.ToCreateNewBook(title)
	if err != nil {
		return errs.BadRequest(err, err.Error())
	}

	err = c.db.Query(ctx, func(r *repository.Repo) error {
		err := r.Books.Create(ctx, book)
		if err != nil {
			return fmt.Errorf("repository Books.Create(): %w", err)
		}
		return nil
	})
	return err
}

func (c BookSrvc) GetByID(ctx context.Context, id entity.BookID) (*entity.Book, error) {
	var book *entity.Book
	err := c.db.Query(ctx, func(r *repository.Repo) error {
		var err error
		book, err = r.Books.GetByID(ctx, id)
		if err != nil {
			return fmt.Errorf("repository Books.GetByID(%d): %w", id, err)
		}
		return nil
	})
	return book, err
}
