package repositories

import (
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"gorm.io/gorm"
)

type StockRepository struct {
	DB *gorm.DB
}

func (s *StockRepository) CreateStock(payload *schemas.UpdateProductStockSchema, productId uint) (*models.Stock, error) {
	newStock := &models.Stock{
		Amount:    payload.Amount,
		ProductID: productId,
	}

	if err := s.DB.Create(newStock).Error; err != nil {
		return nil, err
	}

	return newStock, nil
}
