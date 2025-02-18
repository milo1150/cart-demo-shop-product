package model

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}
