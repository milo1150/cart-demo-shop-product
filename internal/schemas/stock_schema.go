package schemas

type UpdateProductStockPayload struct {
	Amount    uint `json:"amount" validate:"required"`
	ProductId uint `json:"product_id" validate:"required"`
}

type UpdateProductStockSlicesPayload struct {
	Stocks []UpdateProductStockPayload `json:"stocks" validate:"required,dive"`
}
