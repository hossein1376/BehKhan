package maria

import (
	"context"

	"gorm.io/gorm"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
	"github.com/hossein1376/BehKhan/catalogue/pkg/errs"
)

type BooksTable struct {
	db *gorm.DB
}

func newBooksTable(db *gorm.DB) BooksTable {
	return BooksTable{db: db}
}

func (b BooksTable) Create(ctx context.Context, title string) error {
	type books struct {
		Title string `gorm:"title"`
	}
	book := books{
		Title: title,
	}

	r := b.db.Table("books").WithContext(ctx).Create(&book)
	return r.Error
}

func (b BooksTable) GetByID(ctx context.Context, id entity.BookID) (*entity.Book, error) {
	type book struct {
		ID    int64  `gorm:"id"`
		Title string `gorm:"title"`
	}
	var q book
	switch r := b.db.WithContext(ctx).Table("books").Find(&q, id); {
	case r.RowsAffected == 0:
		return nil, errs.NotFound(nil)
	case r.Error != nil:
		return nil, r.Error
	}

	return &entity.Book{ID: entity.BookID(q.ID), Title: q.Title}, nil
}
