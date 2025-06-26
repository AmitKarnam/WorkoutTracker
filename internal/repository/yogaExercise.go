package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type YogaExerciseRepository interface {
	FindAll(ctx context.Context) (*[]models.YogaExercise, error)
	FindByID(ctx context.Context, id uint) (*models.YogaExercise, error)
	FindByName(ctx context.Context, name string) (*models.YogaExercise, error)
	Create(ctx context.Context, strengthExercise *models.YogaExercise) error
	Update(ctx context.Context, strengthExercise *models.YogaExercise) error
	Delete(ctx context.Context, id uint) error
}

type yogaExerciseRepository struct {
	db *gorm.DB
}

func NewYogaExerciseRepository(db *gorm.DB) YogaExerciseRepository {
	return &yogaExerciseRepository{db: db}
}

func (r *yogaExerciseRepository) FindAll(ctx context.Context) (*[]models.YogaExercise, error) {
	return nil, nil
}

func (r *yogaExerciseRepository) FindByID(ctx context.Context, id uint) (*models.YogaExercise, error) {
	return nil, nil
}

func (r *yogaExerciseRepository) FindByName(ctx context.Context, name string) (*models.YogaExercise, error) {
	return nil, nil
}

func (r *yogaExerciseRepository) Create(ctx context.Context, strengthExercise *models.YogaExercise) error {
	return nil
}

func (r *yogaExerciseRepository) Update(ctx context.Context, strengthExercise *models.YogaExercise) error {
	return nil
}

func (r *yogaExerciseRepository) Delete(ctx context.Context, id uint) error {
	return nil
}
