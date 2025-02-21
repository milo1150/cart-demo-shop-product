package schemas

type CreateProductSchema struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description" validate:"max=255"`
	Price       float32 `json:"price" validate:"required"`
	ShopId      uint    `json:"shop_id" validate:"required"`

	// TODO: product_categories
	// TODO: images
}
