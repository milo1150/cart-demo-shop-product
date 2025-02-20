package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Uuid uuid.UUID `json:"uuid" gorm:"not null;type:uuid;unique"`
	Name string    `json:"name" gorm:"not null;unique"`

	Products []Product
	Coupons  []Coupon
}

func (s *Shop) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		s.Uuid = uuidV7
	}
	return nil
}
