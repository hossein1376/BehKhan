package controller

import (
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
)

type Controllers struct {
	BookCtrl BookCtrl
}

func New(db *pool.DB) Controllers {
	return Controllers{
		BookCtrl: newBookCtrl(db),
	}
}
