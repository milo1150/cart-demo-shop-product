package loaders

import (
	"fmt"
	"log"
	"os"
	"shop-product-service/internal/database"
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"github.com/go-faker/faker/v4"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductPgLoader struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func (p *ProductPgLoader) getCreatedShopsJson(shopJson schemas.ShopJsonFile) schemas.CreatedShopsJson {
	shops := schemas.CreatedShopsJson{}

	for _, shopJson := range shopJson.Shops {
		shopModel := models.Shop{}
		if err := p.DB.Where("name = ?", shopJson.Name).Find(&shopModel).Error; err != nil {
			log.Fatalf("ShopJson %v not found: %v", shopJson.Name, err)
		}

		shops[shopJson.TmpShopId] = schemas.CreatedShopJson{
			TmpShopId: shopJson.TmpShopId,
			Name:      shopModel.Name,
			ShopId:    shopModel.ID,
		}
	}

	return shops
}

func (p *ProductPgLoader) GetImageFile(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (p *ProductPgLoader) insertProductToPostgres(product schemas.ProductJson, shops schemas.CreatedShopsJson, minioClient *database.MinIO, bucketName string) {
	// Find shop id (ShopID relation)
	shopId := shops[product.TmpShopId].ShopId

	// Random number for generate price and stock
	randInt, _ := faker.RandomInt(20)

	// Generate image url
	imageUrl := minioClient.GetPublicURL(bucketName, product.Name)

	// Load product from json file into postgres db
	newProduct := models.Product{
		Name:        product.Name,
		Description: product.Name,
		ShopID:      shopId,
		Price:       float32(randInt[0]),
		ImageUrl:    imageUrl,
		Stock:       uint(randInt[1]),
	}

	if err := p.DB.Create(&newProduct).Error; err != nil {
		p.Log.Error(fmt.Sprintf("Failed to create product: %v", product.Name))
	} else {
		p.Log.Info(fmt.Sprintf("Created product: %v", product.Name))
	}
}

func (p *ProductPgLoader) prepareProductsJson(productsJson []schemas.ProductJson, shops schemas.CreatedShopsJson, minioClient *database.MinIO, bucketName string) {
	// Make list of shop json names
	productNames := lo.Map(productsJson, func(productJson schemas.ProductJson, index int) string {
		return productJson.Name
	})

	// Find shops by shopnames
	products := []models.Product{}
	query := p.DB.Where("name IN ?", productNames).Find(&products)
	if query.Error != nil {
		log.Fatalf("InsertShopsJsonToDatabase error: %v", query.Error)
	}

	// If length equal mean all shopsjson already created
	if len(products) == len(productsJson) {
		return
	}

	// Make shopname Hashmap
	existsProducts := map[string]string{}
	for _, product := range products {
		existsProducts[product.Name] = ""
	}

	// Create only product that not in shopname hasmap
	for _, product := range productsJson {
		_, ok := existsProducts[product.Name]
		if !ok {
			p.insertProductToPostgres(product, shops, minioClient, bucketName)
		}
	}
}

func (p *ProductPgLoader) InitializeProductData(minioClient *database.MinIO, bucketName string) {
	// Load & Parse shop.json
	shopJsonFile := LoadShopJsonFile()
	shopJson := ParseShopJsonFile(shopJsonFile)

	// Query created shop.json
	shops := p.getCreatedShopsJson(shopJson)

	// Load & Parse product.json
	productJsonFile := LoadProductJsonFile()
	productsJson := ParseProductJsonFile(productJsonFile)

	// Prepare and create product in product.json into postgres db
	p.prepareProductsJson(productsJson.Products, shops, minioClient, bucketName)
}
