package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Uuid uuid.UUID `json:"uuid" gorm:"not null;type:uuid;uniqueIndex"`
	Name string    `json:"name" gorm:"not null"`

	Products []Product
	Coupons  []Coupon
}
