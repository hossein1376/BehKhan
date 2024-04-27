package service

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/application/service/booksrvc"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/services"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
)

func New(db *pool.DB) services.Service {
	return services.Service{
		BookSrvc: booksrvc.NewBookSrvc(db),
	}
}
