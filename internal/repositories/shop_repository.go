package repositories

import (
	"minicart/internal/models"
	"minicart/internal/schemas"

	"gorm.io/gorm"
)

func CreateShop(db *gorm.DB, payload schemas.CreateShop) error {
	shop := models.Shop{
		Name: payload.Name,
	}

	if err := db.Create(&shop).Error; err != nil {
		return err
	}

	return nil
}
