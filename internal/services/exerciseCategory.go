package services

import (
	"context"
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type ExerciseCategoryService interface {
	GetAll(ctx context.Context) ([]models.ExerciseCategory, error)
	GetByID(ctx context.Context, id uint) (*models.ExerciseCategory, error)
	Create(ctx context.Context, exerciseCategory *models.ExerciseCategory) error
	Delete(ctx context.Context, id uint) error
}

type exerciseCategoryService struct {
	repo repository.ExerciseCategoryRepository
}

func NewExerciseCategoryService(repo repository.ExerciseCategoryRepository) ExerciseCategoryService {
	return &exerciseCategoryService{repo: repo}
}

func (s *exerciseCategoryService) GetAll(ctx context.Context) ([]models.ExerciseCategory, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		return s.repo.FindAll(ctx)
	}
}

func (s *exerciseCategoryService) GetByID(ctx context.Context, id uint) (*models.ExerciseCategory, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		return s.repo.FindByID(ctx, id)
	}
}

func (s *exerciseCategoryService) Create(ctx context.Context, exerciseCategory *models.ExerciseCategory) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		existing, err := s.repo.FindByName(ctx, exerciseCategory.Category)
		if err == nil && existing != nil {
			return fmt.Errorf("exercise category already exists")
		}
		return s.repo.Create(ctx, exerciseCategory)
	}
}

func (s *exerciseCategoryService) Delete(ctx context.Context, id uint) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		_, err := s.repo.FindByID(ctx, id)
		if err != nil {
			return err
		}
		return s.repo.Delete(ctx, id)
	}
}
