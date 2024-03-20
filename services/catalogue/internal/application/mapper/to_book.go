package mapper

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/dto"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entities"
)

func CreateBookRequestToEntity(req *dto.CreateBookRequest) entities.Book { // return err?
	return entities.Book{
		Name: req.Name,
	}
}
