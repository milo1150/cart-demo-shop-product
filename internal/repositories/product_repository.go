package repositories

import (
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (p *ProductRepository) ProductExists(productUuid uuid.UUID) (bool, error) {
	var count int64
	if err := p.DB.Model(&models.Product{}).Where("uuid = ?", productUuid).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (p *ProductRepository) FindProductByUUID(productUuid uuid.UUID) (*models.Product, error) {
	product := &models.Product{}
	if err := p.DB.First(product, "uuid = ?", productUuid).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductRepository) CreateProduct(payload *schemas.CreateProductSchema, uuidV7 uuid.UUID) error {
	newProduct := &models.Product{
		Uuid:        uuidV7,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		ShopID:      payload.ShopId,
	}

	if err := p.DB.Create(newProduct).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateProductStock(productId uint, amount uint) (*models.Product, error) {
	product := &models.Product{}

	if err := p.DB.Model(&models.Product{}).Where("id = ?", productId).UpdateColumn("stock", amount).Error; err != nil {
		return nil, err
	}

	if err := p.DB.First(product, productId).Error; err != nil {
		return nil, err
	}

	return product, nil
}
