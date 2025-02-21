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

	shopExists, err := repositories.ShopExists(p.DB, payload.ShopId)
	if !shopExists || err != nil {
		return errors.New("invalid shop")
	}

	productRepository := repositories.ProductRepository{DB: p.DB}
	if err := productRepository.CreateProduct(payload, uuidV7); err != nil {
		return err
	}

	return nil
}
