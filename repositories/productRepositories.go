package repositories

import (
	"context"
	"gst-billing/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	SearchProduct(ctx context.Context, db *gorm.DB, condition map[string]interface{}) ([]models.Product, error)
	UpdateRecordWithCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}, data map[string]interface{}) (*models.Product, error)
	ReadRecordsWithCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (*[]models.Product, error)
	CreateRecord(ctx context.Context, db *gorm.DB, data *models.Product) error
}

type productRepository struct{}

func NewProductsRepository() *productRepository {
	return &productRepository{}
}

func GetProductRepository() ProductRepository {
	return NewProductsRepository()
}

func (productRepository *productRepository) SearchProduct(ctx context.Context, db *gorm.DB, condition map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	result := db.WithContext(ctx).Where(condition).First(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
func (productRepository *productRepository) UpdateRecordWithCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}, data map[string]interface{}) (*models.Product, error) {
	var product models.Product
	result := db.WithContext(ctx).Model(&product).Where(condition).Updates(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (productRepository *productRepository) CreateRecord(ctx context.Context, db *gorm.DB, data *models.Product) error {
	result := db.WithContext(ctx).Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (productRepository *productRepository) ReadRecordsWithCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (*[]models.Product, error) {
	var product []models.Product
	result := db.WithContext(ctx).Where(condition).Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
