package loader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"shop-product-service/internal/database"
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"github.com/minio/minio-go/v7"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductLoader struct {
	Client *minio.Client
	Ctx    context.Context
	Log    *zap.Logger
	DB     *gorm.DB
}

func (p *ProductLoader) LoadJsonFile() []byte {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get basePath")
	}

	filePath := fmt.Sprintf("%v/internal/assets/product.json", basePath)

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read product.json: %v", err)
	}

	return file
}

func (p *ProductLoader) ParseJsonFile(file []byte) schemas.ProductJsonFile {
	productsJson := schemas.ProductJsonFile{}

	err := json.Unmarshal(file, &productsJson)
	if err != nil {
		log.Fatalf("Failed to parse product.json: %v", err)
	}

	return productsJson
}

func (p *ProductLoader) GetFileContentType(path string) string {
	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
	return contentType
}

func (p *ProductLoader) UploadFilesToMinIO(file schemas.ProductJsonFile) {
	products := file.Products
	minio := database.MinIO{Client: p.Client, Context: p.Ctx}

	basePath, err := os.Getwd()
	if err != nil {
		p.Log.Sugar().Fatalf("Failed to get basePath")
	}

	for _, product := range products {
		isFileExists := minio.FileExists("product-image", product.Name)
		if !isFileExists {
			imagePath := fmt.Sprintf("%v/internal/assets/images/%v", basePath, product.ImageName)
			contentType := p.GetFileContentType(imagePath)
			minio.UploadFile("product-image", product.Name, imagePath, contentType, p.Log)
		}
	}
}

func (p *ProductLoader) GetImageFilePath(filename string) string {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get basePath")
	}

	filePath := fmt.Sprintf("%v/internal/assets/images/%v", basePath, filename)

	return filePath
}

func (p *ProductLoader) GetImageFile(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (p *ProductLoader) insertProductToPostgres(product schemas.ProductJson) {
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

func (p *ProductLoader) prepareProductsJson(productsJson []schemas.ProductJson) {
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

func (p *ProductLoader) InitializeProductData() {
	file := p.LoadJsonFile()

	productsJson := p.ParseJsonFile(file)

	// p.UploadFilesToMinIO(productsJson)

	p.prepareProductsJson(productsJson.Products)
}
