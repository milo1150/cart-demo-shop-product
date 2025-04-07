package repositories

import (
	"fmt"
	"shop-product-service/internal/models"
	"shop-product-service/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (p *ProductRepository) VerifyIsProductExistsByID(productId uint) (bool, error) {
	var count int64
	if err := p.DB.Model(&models.Product{}).Where("id = ?", productId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (p *ProductRepository) VerifyIsProductExistsByUUID(productUuid uuid.UUID) (bool, error) {
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

func (p *ProductRepository) FindProductByID(productId uint) (*models.Product, error) {
	product := &models.Product{}
	if err := p.DB.First(product, "id = ?", productId).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductRepository) FindProductsByIDs(productIds []uint64) (*[]models.Product, error) {
	products := &[]models.Product{}
	if err := p.DB.Where("id IN ?", productIds).Find(products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) CreateProduct(payload schemas.CreateProductSchema) error {
	newProduct := &models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		ShopID:      payload.ShopId,
		Stock:       payload.Stock,
	}

	if err := p.DB.Create(newProduct).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateProductStock(productId uint, amount uint) error {
	result := p.DB.Model(&models.Product{}).Where("id = ?", productId).UpdateColumn("stock", amount)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no product found with ID %d to update stock", productId)
	}

	return nil
}

func (p *ProductRepository) GetProducts(payload schemas.GetProducts) (*[]models.Product, error) {
	products := []models.Product{}

	// Ordered
	var ordered string
	if payload.Ordered {
		ordered = "updated_at desc"
	} else {
		ordered = "RANDOM()"
	}

	query := p.DB.Preload("Shop").Order(ordered).Find(&products)

	if query.Error != nil {
		return nil, query.Error
	}

	return &products, nil
}
