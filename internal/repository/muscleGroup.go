package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type MuscleGroupRepository interface {
	FindAll(ctx context.Context) ([]models.MuscleGroup, error)
	FindByID(ctx context.Context, id uint) (*models.MuscleGroup, error)
	FindByName(ctx context.Context, name string) (*models.MuscleGroup, error)
	Create(ctx context.Context, muscleGroup *models.MuscleGroup) error
	Update(ctx context.Context, muscleGroup *models.MuscleGroup) error
	Delete(ctx context.Context, id uint) error
}

type muscleGroupRepository struct {
	db *gorm.DB
}

func NewMuscleGroupRepository(db *gorm.DB) MuscleGroupRepository {
	return &muscleGroupRepository{db: db}
}

func (r *muscleGroupRepository) FindAll(ctx context.Context) ([]models.MuscleGroup, error) {
	var muscleGroups []models.MuscleGroup
	err := r.db.WithContext(ctx).Find(&muscleGroups).Error
	return muscleGroups, err
}

func (r *muscleGroupRepository) FindByID(ctx context.Context, id uint) (*models.MuscleGroup, error) {
	var muscleGroup models.MuscleGroup
	err := r.db.WithContext(ctx).First(&muscleGroup, id).Error
	return &muscleGroup, err
}

func (r *muscleGroupRepository) FindByName(ctx context.Context, name string) (*models.MuscleGroup, error) {
	var muscleGroup models.MuscleGroup
	err := r.db.WithContext(ctx).Where("LOWER(muscle_group) = LOWER(?)", name).First(&muscleGroup).Error
	return &muscleGroup, err
}

func (r *muscleGroupRepository) Create(ctx context.Context, muscleGroup *models.MuscleGroup) error {
	return r.db.WithContext(ctx).Create(muscleGroup).Error
}

func (r *muscleGroupRepository) Update(ctx context.Context, muscleGroup *models.MuscleGroup) error {
	return r.db.WithContext(ctx).Save(muscleGroup).Error
}

func (r *muscleGroupRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MuscleGroup{}, id).Error
}
