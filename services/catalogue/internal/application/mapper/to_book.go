package mapper

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
)

func ToCreateNewBook(title string) (entity.Book, error) {
	return entity.Book{
		Title: title,
	}, nil
}
