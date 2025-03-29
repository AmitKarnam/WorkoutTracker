package services

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type YogaExerciseService interface {
	GetAll(ctx context.Context) (*[]models.YogaExercise, error)
	GetByID(ctx context.Context, id uint) (*models.YogaExercise, error)
	Create(ctx context.Context, strengthExercise *models.YogaExercise) error
	Update(ctx context.Context, id uint, input models.YogaExercise) error
	Delete(ctx context.Context, id uint) error
}

type yogaExerciseService struct {
	repo *repository.StrengthExerciseRepository
}

func NewYogaExerciseService(repo *repository.StrengthExerciseRepository) YogaExerciseService {
	return &yogaExerciseService{repo: repo}
}

func (s *yogaExerciseService) GetAll(ctx context.Context) (*[]models.YogaExercise, error) {
	return nil, nil
}

func (s *yogaExerciseService) GetByID(ctx context.Context, id uint) (*models.YogaExercise, error) {
	return nil, nil
}

func (s *yogaExerciseService) Create(ctx context.Context, strengthExercise *models.YogaExercise) error {
	return nil
}

func (s *yogaExerciseService) Update(ctx context.Context, id uint, yogaExerciseService models.YogaExercise) error {
	return nil
}

func (s *yogaExerciseService) Delete(ctx context.Context, id uint) error {
	return nil
}
