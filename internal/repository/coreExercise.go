package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type CoreExerciseRepository interface {
	FindAll(ctx context.Context) (*[]models.CoreExercise, error)
	FindByID(ctx context.Context, id uint) (*models.CoreExercise, error)
	FindByName(ctx context.Context, name string) (*models.CoreExercise, error)
	Create(ctx context.Context, strengthExercise *models.CoreExercise) error
	Update(ctx context.Context, strengthExercise *models.CoreExercise) error
	Delete(ctx context.Context, id uint) error
}

type coreExerciseRepository struct {
	db *gorm.DB
}

func NewCoreExerciseRepository(db *gorm.DB) CoreExerciseRepository {
	return &coreExerciseRepository{db: db}
}

func (r *coreExerciseRepository) FindAll(ctx context.Context) (*[]models.CoreExercise, error) {
	return nil, nil
}

func (r *coreExerciseRepository) FindByID(ctx context.Context, id uint) (*models.CoreExercise, error) {
	return nil, nil
}

func (r *coreExerciseRepository) FindByName(ctx context.Context, name string) (*models.CoreExercise, error) {
	return nil, nil
}

func (r *coreExerciseRepository) Create(ctx context.Context, strengthExercise *models.CoreExercise) error {
	return nil
}

func (r *coreExerciseRepository) Update(ctx context.Context, strengthExercise *models.CoreExercise) error {
	return nil
}

func (r *coreExerciseRepository) Delete(ctx context.Context, id uint) error {
	return nil
}
