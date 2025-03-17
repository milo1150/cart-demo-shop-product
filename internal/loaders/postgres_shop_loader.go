package loaders

import (
	"context"
	"fmt"
	"log"
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ShopPgLoader struct {
	Ctx context.Context
	Log *zap.Logger
	DB  *gorm.DB
}

func (s *ShopPgLoader) InsertShopsJsonToDatabase(shopsJson []schemas.ShopJson) {
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

func (s *ShopPgLoader) InitializeShopData() {
	file := LoadShopJsonFile()

	shopsJson := ParseShopJsonFile(file)

	s.InsertShopsJsonToDatabase(shopsJson.Shops)
}
