package repository

import (
	"context"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"gorm.io/gorm"
)

type StrengthExerciseRepository interface {
	FindAll(ctx context.Context) (*[]models.StrengthExercise, error)
	FindByMuscleGroup(ctx context.Context, muscleGroup string) (*[]models.StrengthExercise, error)
	FindByID(ctx context.Context, id uint) (*models.StrengthExercise, error)
	FindByName(ctx context.Context, name string) (*models.StrengthExercise, error)
	Create(ctx context.Context, strengthExercise *models.StrengthExercise) error
	Update(ctx context.Context, strengthExercise *models.StrengthExercise) error
	Delete(ctx context.Context, id uint) error
}

type strengthExerciseRepository struct {
	db                    *gorm.DB
	muscleGroupRepository MuscleGroupRepository
}

func NewStrengthExerciseRepository(db *gorm.DB, muscleGroupRepository MuscleGroupRepository) StrengthExerciseRepository {
	return &strengthExerciseRepository{db: db, muscleGroupRepository: muscleGroupRepository}
}

func (r *strengthExerciseRepository) FindAll(ctx context.Context) (*[]models.StrengthExercise, error) {
	var strengthExercises []models.StrengthExercise
	err := r.db.WithContext(ctx).Find(&strengthExercises).Error
	return &strengthExercises, err
}

func (r *strengthExerciseRepository) FindByMuscleGroup(ctx context.Context, muscleGroupName string) (*[]models.StrengthExercise, error) {
	muscleGroup, err := r.muscleGroupRepository.FindByName(ctx, muscleGroupName)
	if err != nil {
		return nil, err
	}
	muscleGroupID := muscleGroup.ID
	var strengthExercises *[]models.StrengthExercise
	err = r.db.WithContext(ctx).Where("muscle_group_id = ?", muscleGroupID).Find(&strengthExercises).Error
	return strengthExercises, err

}

func (r *strengthExerciseRepository) FindByID(ctx context.Context, id uint) (*models.StrengthExercise, error) {
	var strengthExercise models.StrengthExercise
	err := r.db.WithContext(ctx).Find(&strengthExercise, id).Error
	return &strengthExercise, err
}

func (r *strengthExerciseRepository) FindByName(ctx context.Context, name string) (*models.StrengthExercise, error) {
	var strengthExercise models.StrengthExercise
	err := r.db.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&strengthExercise).Error
	return &strengthExercise, err
}

func (r *strengthExerciseRepository) Create(ctx context.Context, strengthExercise *models.StrengthExercise) error {
	return r.db.WithContext(ctx).Create(strengthExercise).Error
}

func (r *strengthExerciseRepository) Update(ctx context.Context, strengthExercise *models.StrengthExercise) error {
	return r.db.WithContext(ctx).Save(strengthExercise).Error
}

func (r *strengthExerciseRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(models.StrengthExercise{}, id).Error
}
