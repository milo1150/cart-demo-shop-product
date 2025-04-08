package dto

import (
	"shop-product-service/internal/models"
	"time"
)

type ProductDTO struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Stock       uint      `json:"stock"`
	Image       string    `json:"image"`
}

func TransformProductDTO(productModel *models.Product) ProductDTO {
	product := ProductDTO{
		Id:          productModel.ID,
		CreatedAt:   productModel.CreatedAt,
		UpdatedAt:   productModel.UpdatedAt,
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
		Stock:       productModel.Stock,
		Image:       productModel.ImageUrl,
	}
	return product
}

func TransformProductListDTO(productModels *[]models.Product) []ProductDTO {
	products := make([]ProductDTO, len(*productModels))

	for index, product := range *productModels {
		products[index] = TransformProductDTO(&product)
	}

	return products
}
