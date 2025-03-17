package loaders

import (
	"context"
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

type ProductMinIOLoader struct {
	Client *minio.Client
	Ctx    context.Context
	Log    *zap.Logger
}

func (p *ProductMinIOLoader) GetFileContentType(path string) string {
	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
	return contentType
}

func (p *ProductMinIOLoader) UploadFilesToMinIO(file schemas.ProductJsonFile) {
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

func (p *ProductMinIOLoader) GetImageFilePath(filename string) string {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get basePath")
	}

	filePath := fmt.Sprintf("%v/internal/assets/images/%v", basePath, filename)

	return filePath
}

func (p *ProductMinIOLoader) InitializeProductData() {
	file := LoadProductJsonFile()

	productsJson := ParseProductJsonFile(file)

	p.UploadFilesToMinIO(productsJson)
}
