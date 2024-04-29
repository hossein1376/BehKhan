package booksrvc

import (
	"context"
	"fmt"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entities"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
)

type BookSrvc struct {
	db *pool.DB
}

func NewBookSrvc(db *pool.DB) BookSrvc {
	return BookSrvc{db: db}
}

func (c BookSrvc) Create(ctx context.Context, title string) error {
	err := c.db.Query(ctx, func(ctx context.Context, p *pool.Pool) error {
		err := p.Books.Create(ctx, entities.Book{})
		if err != nil {
			return fmt.Errorf("repository Books.Create(): %w", err)
		}
		return nil
	})
	return err
}

func (c BookSrvc) GetByID(ctx context.Context, id entities.BookID) (*entities.Book, error) {
	var book *entities.Book
	err := c.db.Query(ctx, func(ctx context.Context, p *pool.Pool) error {
		var err error
		book, err = p.Books.GetByID(ctx, id)
		if err != nil {
			return fmt.Errorf("repository Books.GetByID(%d): %w", id, err)
		}
		return nil
	})
	return book, err
}
