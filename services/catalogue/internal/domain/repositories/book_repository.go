package repositories

import (
	"context"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entities"
)

type BookRepository interface {
	Create(ctx context.Context, book entities.Book) error
	GetByID(ctx context.Context, id entities.BookID) (*entities.Book, error)
}
