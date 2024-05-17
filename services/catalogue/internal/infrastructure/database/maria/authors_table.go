package maria

import (
	"gorm.io/gorm"
)

type AuthorsTable struct {
	db *gorm.DB
}

func newAuthorsTable(db *gorm.DB) AuthorsTable {
	return AuthorsTable{db: db}
}
