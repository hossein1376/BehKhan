package repository

import (
	"context"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
)

type BookRepository interface {
	Create(ctx context.Context, title string) error
	GetByID(ctx context.Context, id entity.BookID) (*entity.Book, error)
}
