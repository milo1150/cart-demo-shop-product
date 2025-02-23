package services

import (
	"fmt"
	"shop-product-service/internal/dto"
	"shop-product-service/internal/repositories"
	"shop-product-service/internal/schemas"

	"gorm.io/gorm"
)

type StockService struct {
	DB *gorm.DB
}

func (s *StockService) UpdateProductStockService(payload *schemas.UpdateProductStockSchema) (*dto.ProductDTO, error) {
	// Begin transaction with panic handler
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Init repositories DB with tx
	productRepository := repositories.ProductRepository{DB: tx}
	stockRepository := repositories.StockRepository{DB: tx}

	// Find product by uuid
	product, err := productRepository.FindProductByUUID(payload.ProductUuid)
	if err != nil {
		return nil, fmt.Errorf("product not found or invalid uuid: %w", err)
	}

	// Calculate new product stock
	newProductStock := product.Stock + payload.Amount

	// Update Product stock
	updatedProduct, err := productRepository.UpdateProductStock(product.ID, newProductStock)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error update product stock: %w", err)
	}

	// Create Stock log
	_, err = stockRepository.CreateStock(payload, product.ID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error create stock: %w", err)
	}

	// Transform response product
	updatedProductDTO := dto.TransformProductDTO(updatedProduct)

	// Commit UpdateProductStock transaction
	tx.Commit()

	return &updatedProductDTO, nil
}
