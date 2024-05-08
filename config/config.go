package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// TODO: Initializar PostgreSQL
	db, err = InitializerPostgreSQL()
	if err != nil {
		return fmt.Errorf("error initializer postgre: %v", err)
	}

	return nil
}

func GetPostgreSQL() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// TODO: Initializer logger
	logger = NewLogger(prefix)
	return logger
}
