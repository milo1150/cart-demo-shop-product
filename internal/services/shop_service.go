package services

import (
	"errors"
	"minicart/internal/models"
	"minicart/internal/repositories"
	"minicart/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShopService struct {
	DB *gorm.DB
}

func (s *ShopService) CreateShop(payload *schemas.CreateShop) error {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	return repositories.CreateShop(s.DB, payload, uuidV7)
}

func (s *ShopService) GetShopDetail(shopUuid uuid.UUID) (*models.Shop, error) {
	shop, err := repositories.GetShopDetail(s.DB, shopUuid)
	if err != nil {
		return nil, errors.New("shop not found")
	}
	return shop, nil
}
