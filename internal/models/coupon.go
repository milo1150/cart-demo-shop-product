package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Uuid                uuid.UUID `json:"uuid" gorm:"not null;type:uuid"`
	Description         string    `json:"description"`
	Discount            uint32    `json:"discount"`
	DiscountUnit        uint32    `json:"discount_unit"`
	DiscountTrigger     uint32    `json:"discount_trigger"`
	DiscountTriggerUnit uint32    `json:"discount_trigger_unit"`

	ShopID uint
}
