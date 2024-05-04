package mapper

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entities"
)

func ToCreateNewBook(title string) (entities.Book, error) {
	return entities.Book{
		Title: title,
	}, nil
}
