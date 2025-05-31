package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AmitKarnam/WorkoutTracker/internal/auth/jwt"
	"github.com/AmitKarnam/WorkoutTracker/internal/constants"
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
	"gorm.io/gorm"
)

type RefreshTokenService interface {
	CreateOrUpdate(ctx context.Context, userID uint) (string, error)
	GetByToken(ctx context.Context, token string) (*models.RefreshToken, error)
	HandleRefresh(ctx context.Context, token string) (string, string, error)
	Delete(ctx context.Context, token string) error
}

type refreshTokenService struct {
	tokenRepo repository.RefreshTokenRepository
}

func NewRefreshTokenService(tokenRepo repository.RefreshTokenRepository) RefreshTokenService {
	return &refreshTokenService{tokenRepo: tokenRepo}
}

func (s *refreshTokenService) CreateOrUpdate(ctx context.Context, userID uint) (string, error) {
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		// create a new refresh token
		refreshToken, err := jwt.GenerateRefreshToken()
		if err != nil {
			return "", err
		}

		// set expiry time for the refresh token generated
		expiryTime := time.Now().Add(constants.RefreshTokenExpiry)

		// check if an exsisting record already exsists for a customer
		refreshTokenInstance, err := s.tokenRepo.FindByUserID(ctx, userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}

		if refreshTokenInstance == nil {
			// Create a new refresh token instance
			refreshTokenInstance = &models.RefreshToken{
				UserID:     userID,
				Token:      refreshToken,
				ExpiryTime: expiryTime,
			}
		} else {
			// Update the existing refresh token instance
			refreshTokenInstance.Token = refreshToken
			refreshTokenInstance.ExpiryTime = expiryTime
		}

		err = s.tokenRepo.CreateOrUpdate(ctx, refreshTokenInstance)
		if err != nil {
			return "", err
		}
		return refreshToken, nil
	}
}

func (s *refreshTokenService) GetByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		return s.tokenRepo.FindByToken(ctx, token)
	}
}

func (s *refreshTokenService) HandleRefresh(ctx context.Context, token string) (string, string, error) {
	// Check the expiry of the refresh token
	refreshTokenInstance, err := s.tokenRepo.FindByToken(ctx, token)
	if err != nil {
		return "", "", fmt.Errorf("unable to find refresh token instance with given token", err)
	}
	if time.Now().After(refreshTokenInstance.ExpiryTime) {
		return "", "", errors.New("refresh token has expired")
	}

	// Generate new access token, refresh token
	newAccessToken, err := jwt.GenerateUserAccessToken(&refreshTokenInstance.User)
	if err != nil {
		return "", "", fmt.Errorf("error generating user access token", err)
	}

	newRefreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return "", "", fmt.Errorf("error generating refresh token", err)
	}

	// set expiry time for the refresh token generated
	newExpiryTime := time.Now().Add(constants.RefreshTokenExpiry)

	refreshTokenInstance.Token = newRefreshToken
	refreshTokenInstance.ExpiryTime = newExpiryTime

	// update the refresh_token table
	err = s.tokenRepo.CreateOrUpdate(ctx, refreshTokenInstance)
	if err != nil {
		return "", "", fmt.Errorf("error updating database with refrsh token", err)
	}

	// Send back the access token, refresh token
	return newAccessToken, newRefreshToken, nil
}

func (s *refreshTokenService) Delete(ctx context.Context, token string) error {
	tokenInstance, err := s.tokenRepo.FindByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("unable to fetch refresh token")
	}
	return s.tokenRepo.Delete(ctx, tokenInstance.ID)
}
