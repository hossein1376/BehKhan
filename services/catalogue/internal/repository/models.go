package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Book BookRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Book: &BooksTable{db: db},
	}
}
