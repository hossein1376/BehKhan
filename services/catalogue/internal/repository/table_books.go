package repository

import (
	"gorm.io/gorm"

	"github.com/hossein1376/BehKhan/catalogue/internal/transfer"
)

type Book struct {
	gorm.Model
	Name string
}

type BooksTable struct {
	db *gorm.DB
}

func (b *BooksTable) GetAll() ([]Book, error) {
	var books []Book

	res := b.db.Find(&books)
	if res.Error != nil {
		return nil, transfer.InternalError{Err: res.Error}
	}

	return books, nil
}

func (b *BooksTable) GetByID(id int) (*Book, error) {
	var book Book

	res := b.db.Find(&book, id)
	if res.Error != nil {
		return nil, transfer.InternalError{Err: res.Error}
	}

	if res.RowsAffected == 0 {
		return nil, transfer.NotFoundError{}
	}

	return &book, nil
}
