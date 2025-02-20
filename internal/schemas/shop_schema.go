package schemas

type CreateShop struct {
	ShopName string `json:"shop_name" validate:"required"`
}
