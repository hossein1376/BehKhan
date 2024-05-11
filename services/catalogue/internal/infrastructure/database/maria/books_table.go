package maria

import (
	"context"
	"database/sql"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
)

type BooksTable struct {
	tx *sql.Tx
}

func newBooksTable(tx *sql.Tx) BooksTable {
	return BooksTable{tx: tx}
}

func (b BooksTable) Create(ctx context.Context, book entity.Book) error {
	panic("implement me")
}

func (b BooksTable) GetByID(ctx context.Context, id entity.BookID) (*entity.Book, error) {
	panic("implement me")
}
