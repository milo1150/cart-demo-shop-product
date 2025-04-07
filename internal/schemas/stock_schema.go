package schemas

type UpdateProductStockSchema struct {
	Amount    uint `json:"amount" validate:"required"`
	ProductId uint `json:"product_id" validate:"required"`
}
