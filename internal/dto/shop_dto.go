package dto

import (
	"minicart/internal/models"
	"time"

	"github.com/google/uuid"
)

type ShopDTO struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Products  []ProductDTO
}

func TransformShopDTO(shopModel *models.Shop) *ShopDTO {
	shop := &ShopDTO{
		CreatedAt: shopModel.CreatedAt,
		UpdatedAt: shopModel.UpdatedAt,
		Uuid:      shopModel.Uuid,
		Name:      shopModel.Name,
		Products:  TransformProductListDTO(shopModel.Products),
	}
	return shop
}
