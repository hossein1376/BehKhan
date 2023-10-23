package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hossein1376/BehKhan/catalogue/internal/repository"
	"github.com/hossein1376/BehKhan/catalogue/pkg/configs"
)

func GetDB(settings *configs.Settings) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DB.Username,
		settings.DB.Password,
		settings.DB.Host,
		settings.DB.Port,
		settings.DB.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	s, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = s.Ping()
	if err != nil {
		return nil, err
	}

	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&repository.Book{},
	)

	return err
}
