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

func (p *ProductLoader) getFileContentType(path string) string {
	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
	return contentType
}

func (p *ProductLoader) uploadFiles(file schemas.ProductJsonFile) {
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
			contentType := p.getFileContentType(imagePath)
			minio.UploadFile("product-image", product.Name, imagePath, contentType, p.Log)
		}
	}
}

func (p *ProductLoader) InsertProductsJsonToDatabase(productsJson []schemas.ProductJson) {
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
			if err := p.DB.Create(&models.Product{Name: product.Name}).Error; err != nil {
				p.Log.Error(fmt.Sprintf("Failed to create product: %v", product.Name))
			} else {
				p.Log.Info(fmt.Sprintf("Created product: %v", product.Name))
			}
		}
	}
}

func (p *ProductLoader) InitializeProductData() {
	file := p.LoadJsonFile()

	productsJson := p.ParseJsonFile(file)

	p.uploadFiles(productsJson)

	p.InsertProductsJsonToDatabase(productsJson.Products)
}
