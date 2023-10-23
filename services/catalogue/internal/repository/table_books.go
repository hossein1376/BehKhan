package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name string
}

type BooksTable struct {
	db *gorm.DB
}

func (b *BooksTable) GetAll() []Book {
	var books []Book

	b.db.Find(&books)
	fmt.Println(books)

	return books
}
