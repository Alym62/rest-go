package handler

import (
	"github.com/Alym62/rest-go/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializerHandler() {
	logger = config.GetLogger("request")
	db = config.GetPostgreSQL()
}
