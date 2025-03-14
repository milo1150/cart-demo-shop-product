package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Uuid        uuid.UUID `json:"uuid" gorm:"not null;type:uuid;uniqueIndex"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float32   `json:"price" gorm:"not null"`
	Stock       uint      `json:"stock" gorm:"not null"`

	// ProductCategory []ProductCategory `gorm:"many2many:product_categories"`
	ShopID uint `gorm:"not null"`

	// TODO: Image
	// ImageUrl    string `json:"image_url"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		p.Uuid = uuidV7
	}
	return nil
}
