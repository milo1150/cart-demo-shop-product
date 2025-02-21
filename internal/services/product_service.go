package services

import (
	"errors"
	"minicart/internal/repositories"
	"minicart/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (p *ProductService) CreateProduct(payload *schemas.CreateProductSchema) error {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	if err := repositories.FindShop(p.DB, payload.ShopId); err != nil {
		return errors.New("invalid shop")
	}

	if err := repositories.CreateProduct(p.DB, payload, uuidV7); err != nil {
		return err
	}

	return nil
}
