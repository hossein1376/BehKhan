package services

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/application/services/booksrvc"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/service"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
)

func New(db *pool.Pool) service.Service {
	return service.Service{
		BookSrvc: booksrvc.NewBookSrvc(db),
	}
}
