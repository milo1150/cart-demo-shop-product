package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Uuid        uuid.UUID      `json:"uuid" gorm:"not null;type:uuid;uniqueIndex"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float32        `json:"price" gorm:"not null"`
	Stock       uint           `json:"stock" gorm:"not null"`
	ImageUrl    string         `json:"image_url"`
	Image       []byte         `json:"-" gorm:"type:bytea"`

	// Relation
	ShopID uint `json:"shop_id" gorm:"default:null;constraint:OnDelete:SET NULL"`

	//TODO: ProductCategory []ProductCategory `gorm:"many2many:product_categories"`

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
