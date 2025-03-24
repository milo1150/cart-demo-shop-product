package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shop struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Uuid      uuid.UUID      `json:"uuid" gorm:"not null;type:uuid;unique"`
	Name      string         `json:"name" gorm:"not null;unique"`

	// Relation
	Products []Product `json:"products"`
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
