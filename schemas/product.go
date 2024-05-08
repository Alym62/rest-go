package schemas

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Company     string
	Location    string
	Price       float64
	Description string
}
