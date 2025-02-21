package schemas

import "github.com/google/uuid"

type UpdateProductStockSchema struct {
	Amount      uint      `json:"amount" validate:"required"`
	ProductUuid uuid.UUID `json:"product_uuid" validate:"required,uuid"`
	UserInfo    User      `json:"user_info"` // TODO: validate: required
}
