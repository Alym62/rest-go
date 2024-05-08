package method

import (
	"github.com/Alym62/rest-go/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializerHandlerLogger() {
	logger = config.GetLogger("method")
	db = config.GetPostgreSQL()
}
