package schemas

import (
	"time"

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

type ProductResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}
