package services

import (
	"context"
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type UserService interface {
	FindByID(ctx context.Context, id uint) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindByID(ctx context.Context, id uint) (*models.User, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		return s.repo.FindByID(ctx, id)
	}
}
