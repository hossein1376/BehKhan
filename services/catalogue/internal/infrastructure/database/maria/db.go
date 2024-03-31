package maria

import (
	"database/sql"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repositories"
)

type Tables struct {
	Books   repositories.BookRepository
	Authors repositories.AuthorRepository
}

func New(db *sql.Tx) *Tables {
	return &Tables{
		Books:   newBooksTable(db),
		Authors: newAuthorsTable(db),
	}
}
