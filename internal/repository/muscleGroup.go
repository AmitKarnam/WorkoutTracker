package repository

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type MuscleGroupRepository interface {
	FindAll() ([]models.MuscleGroup, error)
	FindByID(id uint) (*models.MuscleGroup, error)
	FindByName(name string) (*models.MuscleGroup, error)
	Create(muscleGroup *models.MuscleGroup) error
	Update(muscleGroup *models.MuscleGroup) error
	Delete(id uint) error
}

type muscleGroupRepository struct {
	db *gorm.DB
}

func NewMuscleGroupRepository(db *gorm.DB) MuscleGroupRepository {
	return &muscleGroupRepository{db: db}
}

func (r *muscleGroupRepository) FindAll() ([]models.MuscleGroup, error) {
	var muscleGroups []models.MuscleGroup
	err := r.db.Find(&muscleGroups).Error
	return muscleGroups, err
}

func (r *muscleGroupRepository) FindByID(id uint) (*models.MuscleGroup, error) {
	var muscleGroup models.MuscleGroup
	err := r.db.First(&muscleGroup, id).Error
	return &muscleGroup, err
}

func (r *muscleGroupRepository) FindByName(name string) (*models.MuscleGroup, error) {
	var muscleGroup models.MuscleGroup
	err := r.db.Where("LOWER(muscle_group) = LOWER(?)", name).First(&muscleGroup).Error
	return &muscleGroup, err
}

func (r *muscleGroupRepository) Create(muscleGroup *models.MuscleGroup) error {
	return r.db.Create(muscleGroup).Error
}

func (r *muscleGroupRepository) Update(muscleGroup *models.MuscleGroup) error {
	return r.db.Save(muscleGroup).Error
}

func (r *muscleGroupRepository) Delete(id uint) error {
	return r.db.Delete(&models.MuscleGroup{}, id).Error
}
