package config

import (
	"github.com/Alym62/rest-go/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializerPostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgre")

	// TODO: Create conection DB
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=go port=5432"), &gorm.Config{})
	if err != nil {
		logger.Errf("PostgreSQL error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Product{})
	if err != nil {
		logger.Errf("PostgreSQL automigrate: %v", err)
		return nil, err
	}

	return db, nil
}
