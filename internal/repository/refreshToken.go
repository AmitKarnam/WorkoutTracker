package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	CreateOrUpdate(ctx context.Context, refreshToken *models.RefreshToken) error
	FindByUserID(ctx context.Context, userID uint) (*models.RefreshToken, error)
	FindByToken(ctx context.Context, token string) (*models.RefreshToken, error)
}

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{db: db}
}

func (r *refreshTokenRepository) CreateOrUpdate(ctx context.Context, refreshToken *models.RefreshToken) error {
	return r.db.WithContext(ctx).Save(refreshToken).Error
}

func (r *refreshTokenRepository) FindByUserID(ctx context.Context, userID uint) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&refreshToken).Error
	return &refreshToken, err
}

func (r *refreshTokenRepository) FindByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := r.db.WithContext(ctx).Where("token = ?", token).Find(&refreshToken).Error
	return &refreshToken, err
}
