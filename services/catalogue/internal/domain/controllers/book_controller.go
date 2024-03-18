package controllers

import (
	"context"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/dto"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entities"
)

type BookController interface {
	Create(ctx context.Context, request *dto.CreateBookRequest) error
	GetByID(ctx context.Context, request *dto.GetBookByIDRequest) (entities.Book, error)
}
