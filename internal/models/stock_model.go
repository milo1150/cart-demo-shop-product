package models

import (
	"shop-product-service/internal/schemas"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Amount    uint           `json:"amount"`
	UserInfo  schemas.User   `json:"user_info" gorm:"type:jsonb"` // TODO: not null

	ProductID uint `gorm:"not null"`
	Product   Product
}
