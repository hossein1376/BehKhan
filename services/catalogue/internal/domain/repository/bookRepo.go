package repository

import (
	"context"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
)

type BookRepository interface {
	Create(ctx context.Context, book entity.Book) error
	GetByID(ctx context.Context, id int) (*entity.Book, error)
}
