package loader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ShopLoader struct {
	Ctx context.Context
	Log *zap.Logger
	DB  *gorm.DB
}

func (s *ShopLoader) LoadJsonFile() []byte {
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

func (s *ShopLoader) ParseJsonFile(file []byte) schemas.ShopJsonFile {
	shopsJson := schemas.ShopJsonFile{}

	if err := json.Unmarshal(file, &shopsJson); err != nil {
		log.Fatalf("Failed to parse shop.json: %v", err)
	}

	return shopsJson
}

func (s *ShopLoader) InsertShopsJsonToDatabase(shopsJson []schemas.ShopJson) {
	// Make list of shop json names
	shopNames := lo.Map(shopsJson, func(shopJson schemas.ShopJson, index int) string {
		return shopJson.Name
	})

	// Find shops by shopnames
	shops := []models.Shop{}
	query := s.DB.Where("name IN ?", shopNames).Find(&shops)
	if query.Error != nil {
		log.Fatalf("InsertShopsJsonToDatabase error: %v", query.Error)
	}

	// If length equal mean all shopsjson already created
	if len(shops) == len(shopsJson) {
		return
	}

	// Make shopname Hashmap
	existsShops := map[string]string{}
	for _, shop := range shops {
		existsShops[shop.Name] = ""
	}

	// Create only shop that not in shopname hasmap
	for _, shop := range shopsJson {
		_, ok := existsShops[shop.Name]
		if !ok {
			if err := s.DB.Create(&models.Shop{Name: shop.Name}).Error; err != nil {
				s.Log.Error(fmt.Sprintf("Failed to create shop: %v", shop.Name))
			} else {
				s.Log.Info(fmt.Sprintf("Created shop: %v", shop.Name))
			}
		}
	}
}

func (s *ShopLoader) InitializeShopData() {
	file := s.LoadJsonFile()

	shopsJson := s.ParseJsonFile(file)

	s.InsertShopsJsonToDatabase(shopsJson.Shops)
}
