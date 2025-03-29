package services

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type CoreExerciseService interface {
	GetAll(ctx context.Context) (*[]models.CoreExercise, error)
	GetByID(ctx context.Context, id uint) (*models.CoreExercise, error)
	Create(ctx context.Context, strengthExercise *models.CoreExercise) error
	Update(ctx context.Context, id uint, input models.CoreExercise) error
	Delete(ctx context.Context, id uint) error
}

type coreExerciseService struct {
	repo *repository.StrengthExerciseRepository
}

func NewCoreExerciseService(repo *repository.StrengthExerciseRepository) CoreExerciseService {
	return &coreExerciseService{repo: repo}
}

func (s *coreExerciseService) GetAll(ctx context.Context) (*[]models.CoreExercise, error) {
	return nil, nil
}

func (s *coreExerciseService) GetByID(ctx context.Context, id uint) (*models.CoreExercise, error) {
	return nil, nil
}

func (s *coreExerciseService) Create(ctx context.Context, strengthExercise *models.CoreExercise) error {
	return nil
}

func (s *coreExerciseService) Update(ctx context.Context, id uint, coreExerciseService models.CoreExercise) error {
	return nil
}

func (s *coreExerciseService) Delete(ctx context.Context, id uint) error {
	return nil
}
