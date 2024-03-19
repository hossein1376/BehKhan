package database

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repositories"
)

type Tables struct {
	Books   repositories.BookRepository
	Authors repositories.AuthorRepository
}
