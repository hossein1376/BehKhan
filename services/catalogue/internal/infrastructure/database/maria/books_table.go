package maria

import (
	"context"

	"gorm.io/gorm"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
)

type BooksTable struct {
	db *gorm.DB
}

func newBooksTable(db *gorm.DB) BooksTable {
	return BooksTable{db: db}
}

func (b BooksTable) Create(ctx context.Context, book entity.Book) error {
	panic("implement me")
}

func (b BooksTable) GetByID(ctx context.Context, id entity.BookID) (*entity.Book, error) {
	panic("implement me")
}
