package models

import (
	"shop-product-service/internal/schemas"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Amount   uint         `json:"amount"`
	UserInfo schemas.User `json:"user_info" gorm:"type:jsonb"` // TODO: not null

	ProductID uint `gorm:"not null"`
	Product   Product
}
