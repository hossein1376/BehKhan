package maria

import (
	"gorm.io/gorm"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repository"
)

func New(db *gorm.DB) repository.Tables {
	return repository.Tables{
		Books:   newBooksTable(db),
		Authors: newAuthorsTable(db),
	}
}
