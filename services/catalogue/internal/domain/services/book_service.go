package services

import (
	"context"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entities"
)

type BookService interface {
	Create(ctx context.Context, title string) error
	GetByID(ctx context.Context, id entities.BookID) (*entities.Book, error)
}
