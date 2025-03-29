package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type StrengthExerciseRepository interface {
	FindAll(ctx context.Context) (*[]models.StrengthExercise, error)
	FindByID(ctx context.Context, id uint) (*models.StrengthExercise, error)
	FindByName(ctx context.Context, name string) (*models.StrengthExercise, error)
	Create(ctx context.Context, strengthExercise *models.StrengthExercise) error
	Update(ctx context.Context, strengthExercise *models.StrengthExercise) error
	Delete(ctx context.Context, id uint) error
}

type strengthExerciseRepository struct {
	db *gorm.DB
}

func NewStrengthExerciseRepository(db *gorm.DB) StrengthExerciseRepository {
	return &strengthExerciseRepository{db: db}
}

func (r *strengthExerciseRepository) FindAll(ctx context.Context) (*[]models.StrengthExercise, error) {
	return nil, nil
}

func (r *strengthExerciseRepository) FindByID(ctx context.Context, id uint) (*models.StrengthExercise, error) {
	return nil, nil
}

func (r *strengthExerciseRepository) FindByName(ctx context.Context, name string) (*models.StrengthExercise, error) {
	return nil, nil
}

func (r *strengthExerciseRepository) Create(ctx context.Context, strengthExercise *models.StrengthExercise) error {
	return nil
}

func (r *strengthExerciseRepository) Update(ctx context.Context, strengthExercise *models.StrengthExercise) error {
	return nil
}

func (r *strengthExerciseRepository) Delete(ctx context.Context, id uint) error {
	return nil
}
