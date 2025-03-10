package services

import (
	"errors"
	"shop-product-service/internal/repositories"
	"shop-product-service/internal/schemas"

	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (p *ProductService) CreateProduct(payload schemas.CreateProductSchema) error {
	shopExists, err := repositories.ShopExists(p.DB, payload.ShopId)
	if !shopExists || err != nil {
		return errors.New("invalid shop")
	}

	productRepository := repositories.ProductRepository{DB: p.DB}
	if err := productRepository.CreateProduct(payload); err != nil {
		return err
	}

	return nil
}
