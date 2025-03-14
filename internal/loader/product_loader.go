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
	"shop-product-service/internal/schemas"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type ProductLoader struct {
	Client *minio.Client
	Ctx    context.Context
	Log    *zap.Logger
}

func (p *ProductLoader) loadJsonFile() []byte {
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

func (p *ProductLoader) parseJsonFile(file []byte) schemas.ProductJsonFile {
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

	basePath, err := os.Getwd()
	if err != nil {
		p.Log.Sugar().Fatalf("Failed to get basePath")
	}

	for _, product := range products {
		imagePath := fmt.Sprintf("%v/internal/assets/images/%v", basePath, product.ImageName)
		contentType := p.getFileContentType(imagePath)
		database.UploadFile(p.Client, p.Ctx, "product-image", product.Name, imagePath, contentType, p.Log)
	}

}

func (p *ProductLoader) InitializeProductDatas() {
	file := p.loadJsonFile()

	productsJson := p.parseJsonFile(file)

	p.uploadFiles(productsJson)
}
