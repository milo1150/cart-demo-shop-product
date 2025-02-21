package dto

import (
	"minicart/internal/models"
	"time"

	"github.com/google/uuid"
)

type ProductDTO struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Stock       uint32    `json:"stock"`
}

func TransformProductDTO(productModel *models.Product) ProductDTO {
	product := ProductDTO{
		CreatedAt:   productModel.CreatedAt,
		UpdatedAt:   productModel.UpdatedAt,
		Uuid:        productModel.Uuid,
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
		Stock:       productModel.Stock,
	}
	return product
}

func TransformProductListDTO(productModels []models.Product) []ProductDTO {
	products := make([]ProductDTO, len(productModels))

	for index, product := range productModels {
		products[index] = TransformProductDTO(&product)
	}

	return products
}
