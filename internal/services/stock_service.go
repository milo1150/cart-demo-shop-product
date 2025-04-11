package services

import (
	"fmt"
	"shop-product-service/internal/dto"
	"shop-product-service/internal/enums"
	"shop-product-service/internal/repositories"
	"shop-product-service/internal/schemas"

	"gorm.io/gorm"
)

type StockService struct {
	DB *gorm.DB
}

func (s *StockService) CalculateNewStock(action enums.StockAction, currentStock, amount uint) uint {
	if action == enums.DecreaseStock && amount > currentStock {
		return 0
	}
	if action == enums.IncreaseStock {
		return currentStock + amount
	}
	if action == enums.DecreaseStock {
		return currentStock - amount
	}
	if action == enums.UpdateStock && amount <= 0 {
		return 0
	}
	if action == enums.UpdateStock {
		return amount
	}
	return 0
}

func (s *StockService) UpdateProductStock(payload *schemas.UpdateProductStockSlicesPayload) (*[]dto.ProductDTO, error) {
	var updatedProducts []dto.ProductDTO

	// Begin transaction with panic handler
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// Init repositories DB with tx
		productRepo := repositories.ProductRepository{DB: tx}
		stockRepo := repositories.StockRepository{DB: tx}

		for _, stockPayload := range payload.Stocks {
			// Find product by id
			product, err := productRepo.FindProductByID(stockPayload.ProductId)
			if err != nil {
				return err
			}

			// Calculate new product stock
			newStock := s.CalculateNewStock(stockPayload.Action, product.Stock, stockPayload.Amount)

			// Update Product stock (DB)
			err = productRepo.UpdateProductStock(product.ID, newStock)
			if err != nil {
				return err
			}

			// Update Product stock (Exists query)
			product.Stock = newStock

			// Create Stock log
			_, err = stockRepo.CreateStock(stockPayload, product.ID)
			if err != nil {
				return fmt.Errorf("failed to create stock: %w", err)
			}

			// Transform response product
			updatedProducts = append(updatedProducts, dto.TransformProductDTO(product))
		}

		// Commit transaction
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &updatedProducts, nil
}
