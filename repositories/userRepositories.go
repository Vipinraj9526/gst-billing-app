package repositories

import (
	"context"
	"gst-billing/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ReadRecordWithCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (*models.User, error)
}

type userRepository struct{}

func NewUsersRepository() *userRepository {
	return &userRepository{}
}

func GetUserRepository() UserRepository {
	return NewUsersRepository()
}

func (userRepository *userRepository) ReadRecordWithCondition(ctx context.Context, db *gorm.DB, condition map[string]interface{}) (*models.User, error) {
	var user models.User
	result := db.WithContext(ctx).Where(condition).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
