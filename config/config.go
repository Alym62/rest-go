package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	// TODO: Initializer logger
	logger = NewLogger(prefix)
	return logger
}
