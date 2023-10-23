package repository

import (
	"gorm.io/gorm"
)

type Models struct {
	Book BookRepository
}

func NewModels(db *gorm.DB) *Models {
	return &Models{
		Book: &BooksTable{db: db},
	}
}
