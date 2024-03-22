package service

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
)

type Service struct {
	BookSrvc BookSrvc
}

func New(db *pool.DB) Service {
	return Service{
		BookSrvc: newBookSrvc(db),
	}
}
