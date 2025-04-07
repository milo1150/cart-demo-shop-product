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

func (s *StockService) UpdateProductStockService(payload *schemas.UpdateProductStockSlicesPayload) (*[]dto.ProductDTO, error) {
	var updatedProductsDTO []dto.ProductDTO

	// Begin transaction with panic handler
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// Init repositories DB with tx
		productRepository := repositories.ProductRepository{DB: tx}
		stockRepository := repositories.StockRepository{DB: tx}

		for _, stockPayload := range payload.Stocks {
			// Find product by id
			product, err := productRepository.FindProductByID(stockPayload.ProductId)
			if err != nil {
				return err
			}

			// Calculate new product stock
			newProductStock := product.Stock + stockPayload.Amount

			// Update Product stock (DB)
			err = productRepository.UpdateProductStock(product.ID, newProductStock)
			if err != nil {
				return err
			}

			// Update Product stock (Exists query)
			product.Stock = newProductStock

			// Create Stock log
			_, err = stockRepository.CreateStock(stockPayload, product.ID)
			if err != nil {
				return fmt.Errorf("failed to create stock: %w", err)
			}

			// Transform response product
			updatedProductsDTO = append(updatedProductsDTO, dto.TransformProductDTO(product))
		}

		// Commit transaction
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &updatedProductsDTO, nil
}
