package controller

import (
	"context"

	"github.com/hossein1376/BehKhan/catalogue/internal/application/mapper"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/dto"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entities"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
)

type BookCtrl struct {
	db *pool.DB
}

func newBookCtrl(db *pool.DB) BookCtrl {
	return BookCtrl{db: db}
}

func (c BookCtrl) Create(ctx context.Context, request *dto.CreateBookRequest) error {
	book := mapper.CreateBookRequestToEntity(request)
	err := c.db.Query(ctx, func(ctx context.Context, p *pool.Pool) error {
		return p.Books.Create(ctx, book)
	})
	return err
}

func (c BookCtrl) GetByID(ctx context.Context, request *dto.GetBookByIDRequest) (entities.Book, error) {
	panic("implement me")
}
