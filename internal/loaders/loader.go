package loaders

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"shop-product-service/internal/schemas"

	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadProductJsonFile() []byte {
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

func ParseProductJsonFile(file []byte) schemas.ProductJsonFile {
	productsJson := schemas.ProductJsonFile{}

	err := json.Unmarshal(file, &productsJson)
	if err != nil {
		log.Fatalf("Failed to parse product.json: %v", err)
	}

	return productsJson
}

func LoadShopJsonFile() []byte {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get basePath")
	}

	filePath := fmt.Sprintf("%v/internal/assets/shop.json", basePath)

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read shop.json: %v", err)
	}

	return file
}

func ParseShopJsonFile(file []byte) schemas.ShopJsonFile {
	shopsJson := schemas.ShopJsonFile{}

	if err := json.Unmarshal(file, &shopsJson); err != nil {
		log.Fatalf("Failed to parse shop.json: %v", err)
	}

	return shopsJson
}

func GetImageFilePath(filename string) string {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get basePath")
	}

	filePath := fmt.Sprintf("%v/internal/assets/images/%v", basePath, filename)

	return filePath
}

// Ex: .webp
func GetFileExtension(path string) string {
	ext := filepath.Ext(path)
	return ext
}

// Ex: image/webp
func GetFileMIMEType(path string) string {
	ext := GetFileExtension(path)
	mimeType := mime.TypeByExtension(ext)
	return mimeType
}
