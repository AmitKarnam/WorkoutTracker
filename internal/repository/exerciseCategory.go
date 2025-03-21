package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type ExerciseCategoryRepository interface {
	FindAll(ctx context.Context) ([]models.ExerciseCategory, error)
	FindByID(ctx context.Context, id uint) (*models.ExerciseCategory, error)
	FindByName(ctx context.Context, name string) (*models.ExerciseCategory, error)
	Create(ctx context.Context, exerciseCategory *models.ExerciseCategory) error
	Delete(ctx context.Context, id uint) error
}

type exerciseCategoryRepository struct {
	db *gorm.DB
}

func NewExerciseCategoryRepository(db *gorm.DB) ExerciseCategoryRepository {
	return &exerciseCategoryRepository{db: db}
}

func (r *exerciseCategoryRepository) FindAll(ctx context.Context) ([]models.ExerciseCategory, error) {
	var exerciseCategories []models.ExerciseCategory
	err := r.db.WithContext(ctx).Find(&exerciseCategories).Error
	return exerciseCategories, err
}

func (r *exerciseCategoryRepository) FindByID(ctx context.Context, id uint) (*models.ExerciseCategory, error) {
	var exerciseCategory models.ExerciseCategory
	err := r.db.WithContext(ctx).First(&exerciseCategory, id).Error
	return &exerciseCategory, err
}

func (r *exerciseCategoryRepository) FindByName(ctx context.Context, name string) (*models.ExerciseCategory, error) {
	var exerciseCategory models.ExerciseCategory
	err := r.db.WithContext(ctx).Where("LOWER(category) = LOWER(?)", name).First(&exerciseCategory).Error
	return &exerciseCategory, err
}

func (r *exerciseCategoryRepository) Create(ctx context.Context, exerciseCategory *models.ExerciseCategory) error {
	return r.db.WithContext(ctx).Create(exerciseCategory).Error
}

func (r *exerciseCategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ExerciseCategory{}, id).Error
}
