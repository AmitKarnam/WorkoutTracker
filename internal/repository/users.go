package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindOrCreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FindByID(ctx context.Context, id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindOrCreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	var existingUser models.User
	err := r.db.WithContext(ctx).Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return &existingUser, nil
	}
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Find(&user, id).Error
	return &user, err
}
