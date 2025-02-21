package repositories

import (
	"minicart/internal/models"
	"minicart/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ShopExists(db *gorm.DB, shopId uint) (bool, error) {
	var count int64

	// More efficient because COUNT(*) is faster than loading a full row.
	// Prevents unnecessary data fetching (no need to load an entire Shop record).
	if err := db.Model(&models.Shop{}).Where("id = ?", shopId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

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

func GetShopDetail(db *gorm.DB, shopUuid uuid.UUID) (*models.Shop, error) {
	shop := &models.Shop{}
	if err := db.Preload("Products").First(shop, "uuid = ?", shopUuid).Error; err != nil {
		return nil, err
	}

	return shop, nil
}
