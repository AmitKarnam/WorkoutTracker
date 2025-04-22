package repository

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindOrCreateUser(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindOrCreateUser(user *models.User) (*models.User, error) {
	var existingUser models.User
	err := r.db.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return &existingUser, nil
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
