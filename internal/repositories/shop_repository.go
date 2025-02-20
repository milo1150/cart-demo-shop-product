package repositories

import (
	"minicart/internal/models"
	"minicart/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateShop(db *gorm.DB, payload *schemas.CreateShop, uuid uuid.UUID) error {
	shop := models.Shop{
		Name: payload.ShopName,
		Uuid: uuid,
	}

	if err := db.Create(&shop).Error; err != nil {
		return err
	}

	return nil
}
