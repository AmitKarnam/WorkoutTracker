package services

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type StrengthExerciseService interface {
	GetAll(ctx context.Context) (*[]models.StrengthExercise, error)
	GetByID(ctx context.Context, id uint) (*models.StrengthExercise, error)
	Create(ctx context.Context, strengthExercise *models.StrengthExercise) error
	Update(ctx context.Context, id uint, input models.StrengthExercise) error
	Delete(ctx context.Context, id uint) error
}

type strengthExerciseService struct {
	repo *repository.StrengthExerciseRepository
}

func NewStrengthExerciseService(repo *repository.StrengthExerciseRepository) StrengthExerciseService {
	return &strengthExerciseService{repo: repo}
}

func (s *strengthExerciseService) GetAll(ctx context.Context) (*[]models.StrengthExercise, error) {
	return nil, nil
}

func (s *strengthExerciseService) GetByID(ctx context.Context, id uint) (*models.StrengthExercise, error) {
	return nil, nil
}

func (s *strengthExerciseService) Create(ctx context.Context, strengthExercise *models.StrengthExercise) error {
	return nil
}

func (s *strengthExerciseService) Update(ctx context.Context, id uint, strengthExerciseService models.StrengthExercise) error {
	return nil
}

func (s *strengthExerciseService) Delete(ctx context.Context, id uint) error {
	return nil
}
