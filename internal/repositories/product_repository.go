package repositories

import (
	"minicart/internal/models"
	"minicart/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, payload *schemas.CreateProductSchema, uuidV7 uuid.UUID) error {
	newProduct := &models.Product{
		Uuid:        uuidV7,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		ShopID:      payload.ShopId,
	}

	if err := db.Create(newProduct).Error; err != nil {
		return err
	}

	return nil
}
