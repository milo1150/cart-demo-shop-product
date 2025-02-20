package services

import (
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
