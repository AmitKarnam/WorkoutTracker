package repository

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type ExerciseCategoryRepository interface {
	FindAll() ([]models.ExerciseCategory, error)
	FindByID(id uint) (*models.ExerciseCategory, error)
	FindByName(name string) (*models.ExerciseCategory, error)
	Create(exerciseCategory *models.ExerciseCategory) error
	Delete(id uint) error
}

type exerciseCategoryRepository struct {
	db *gorm.DB
}

func NewExerciseCategoryRepository(db *gorm.DB) ExerciseCategoryRepository {
	return &exerciseCategoryRepository{db: db}
}

func (r *exerciseCategoryRepository) FindAll() ([]models.ExerciseCategory, error) {
	var exerciseCategories []models.ExerciseCategory
	err := r.db.Find(&exerciseCategories).Error
	return exerciseCategories, err
}

func (r *exerciseCategoryRepository) FindByID(id uint) (*models.ExerciseCategory, error) {
	var exerciseCategory models.ExerciseCategory
	err := r.db.First(&exerciseCategory, id).Error
	return &exerciseCategory, err
}

func (r *exerciseCategoryRepository) FindByName(name string) (*models.ExerciseCategory, error) {
	var exerciseCategory models.ExerciseCategory
	err := r.db.Where("LOWER(category) = LOWER(?)", name).First(&exerciseCategory).Error
	return &exerciseCategory, err
}

func (r *exerciseCategoryRepository) Create(exerciseCategory *models.ExerciseCategory) error {
	return r.db.Create(exerciseCategory).Error
}

func (r *exerciseCategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.ExerciseCategory{}, id).Error
}
