package repository

import (
	"gorm.io/gorm"

	"github.com/hossein1376/BehKhan/catalogue/pkg/transfer"
)

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
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

func (b *BooksTable) GetByID(id ...int64) ([]Book, error) {
	var books []Book

	res := b.db.Find(&books, id)
	if res.Error != nil {
		return nil, transfer.InternalError{Err: res.Error}
	}

	if res.RowsAffected == 0 {
		return nil, transfer.NotFoundError{}
	}

	return books, nil
}
