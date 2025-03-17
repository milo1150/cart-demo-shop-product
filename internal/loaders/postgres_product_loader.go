package loaders

import (
	"context"
	"fmt"
	"log"
	"os"
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductPgLoader struct {
	DB  *gorm.DB
	Ctx context.Context
	Log *zap.Logger
}

func (p *ProductPgLoader) GetImageFilePath(filename string) string {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get basePath")
	}

	filePath := fmt.Sprintf("%v/internal/assets/images/%v", basePath, filename)

	return filePath
}

func (p *ProductPgLoader) GetImageFile(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (p *ProductPgLoader) insertProductToPostgres(product schemas.ProductJson) {
	// Find product image and convert to binary
	filePath := p.GetImageFilePath(product.ImageName)
	file, err := p.GetImageFile(filePath)
	if err != nil {
		p.Log.Error(fmt.Sprintf("Failed to create product: %v", product.Name))
	}

	// Load product from json file into postgres db
	newProduct := models.Product{Name: product.Name, Description: product.Name, Image: file}
	if err := p.DB.Create(&newProduct).Error; err != nil {
		p.Log.Error(fmt.Sprintf("Failed to create product: %v", product.Name))
	} else {
		p.Log.Info(fmt.Sprintf("Created product: %v", product.Name))
	}
}

func (p *ProductPgLoader) prepareProductsJson(productsJson []schemas.ProductJson) {
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
			p.insertProductToPostgres(product)
		}
	}
}

func (p *ProductPgLoader) InitializeProductData() {
	file := LoadProductJsonFile()

	productsJson := ParseProductJsonFile(file)

	p.prepareProductsJson(productsJson.Products)
}
