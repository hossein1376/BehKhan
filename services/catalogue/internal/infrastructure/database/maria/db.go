package maria

import (
	"database/sql"

	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database"
)

func New(db *sql.Tx) *database.Tables {
	return &database.Tables{
		Books:   newBooksTable(db),
		Authors: newAuthorsTable(db),
	}
}
