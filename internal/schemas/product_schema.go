package schemas

type CreateProductSchema struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description" validate:"max=255"`
	Price       float32 `json:"price" validate:"required"`
	ShopId      uint    `json:"shop_id" validate:"required"`
	Stock       uint    `json:"stock" validate:"numeric"`

	// TODO: product_categories
	// TODO: images
}

type GenerateProductSchema struct {
	ShopId uint `json:"shop_id" validate:"required"`
}
