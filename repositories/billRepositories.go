package repositories

import (
	"context"
	"gst-billing/models"

	"gorm.io/gorm"
)

type BillRepository interface {
	CreateRecord(ctx context.Context, db *gorm.DB, data *models.Bill) error
}

type billRepository struct{}

func NewBillsRepository() *billRepository {
	return &billRepository{}
}

func GetBillRepository() BillRepository {
	return NewBillsRepository()
}

func (billRepository *billRepository) CreateRecord(ctx context.Context, db *gorm.DB, data *models.Bill) error {
	result := db.WithContext(ctx).Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
