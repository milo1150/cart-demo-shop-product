package loaders

import (
	"context"
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

func (p *ProductMinIOLoader) UploadFilesToMinIO(file schemas.ProductJsonFile, bucketName string, minioClient *database.MinIO) {
	products := file.Products

	for _, product := range products {
		isFileExists := minioClient.FileExists(bucketName, product.Name)
		if !isFileExists {
			imagePath := GetImageFilePath(product.ImageName)
			mimeType := GetFileMIMEType(imagePath)
			minioClient.UploadFile(bucketName, product.Name, imagePath, mimeType, p.Log)
		}
	}
}

func (p *ProductMinIOLoader) InitializeProductData(bucketName string, minioClient *database.MinIO) {
	file := LoadProductJsonFile()

	productsJson := ParseProductJsonFile(file)

	p.UploadFilesToMinIO(productsJson, bucketName, minioClient)
}
