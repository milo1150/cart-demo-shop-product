package schemas

import "shop-product-service/internal/enums"

type UpdateProductStockPayload struct {
	Amount    uint              `json:"amount" validate:"required"`
	ProductId uint              `json:"product_id" validate:"required"`
	Action    enums.StockAction `json:"action" validate:"required,stock_action"`
}

type UpdateProductStockSlicesPayload struct {
	Stocks []UpdateProductStockPayload `json:"stocks" validate:"required,dive"`
}
