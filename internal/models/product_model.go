package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Uuid  uuid.UUID `json:"uuid" gorm:"not null;type:uuid;uniqueIndex"`
	Name  string    `json:"name" gorm:"not null"`
	Price float32   `json:"price"`
	Stock uint32    `json:"stock"`

	ProductCategory *[]ProductCategory `gorm:"many2many:product_categories"`
	ShopID          uint

	// TODO: Image
}
