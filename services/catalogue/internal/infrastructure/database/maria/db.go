package maria

import (
	"database/sql"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repository"
)

func New(db *sql.Tx) repository.Tables {
	return repository.Tables{
		Books:   newBooksTable(db),
		Authors: newAuthorsTable(db),
	}
}
