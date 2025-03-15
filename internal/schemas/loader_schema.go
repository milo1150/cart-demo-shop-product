package schemas

type ShopJsonFile struct {
	Shops []ShopJson
}

type ShopJson struct {
	TmpShopId uint   `json:"tmp_shop_id"`
	Name      string `json:"name"`
}

type ProductJsonFile struct {
	Products []ProductJson
}

type ProductJson struct {
	TmpShopId uint   `json:"tmp_shop_id"`
	Name      string `json:"name"`
	ImageName string `json:"image_name"`
}
