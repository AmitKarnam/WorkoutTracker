package services

import (
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type ExerciseCategoryService interface {
	GetAll() ([]models.ExerciseCategory, error)
	GetByID(id uint) (*models.ExerciseCategory, error)
	Create(exerciseCategory *models.ExerciseCategory) error
	Delete(id uint) error
}

type exerciseCategoryService struct {
	repo repository.ExerciseCategoryRepository
}

func NewExerciseCategoryService(repo repository.ExerciseCategoryRepository) ExerciseCategoryService {
	return &exerciseCategoryService{repo: repo}
}

func (s *exerciseCategoryService) GetAll() ([]models.ExerciseCategory, error) {
	return s.repo.FindAll()
}

func (s *exerciseCategoryService) GetByID(id uint) (*models.ExerciseCategory, error) {
	return s.repo.FindByID(id)
}

func (s *exerciseCategoryService) Create(exerciseCategory *models.ExerciseCategory) error {
	existing, err := s.repo.FindByName(exerciseCategory.Category)
	if err == nil && existing != nil {
		return fmt.Errorf("exercise category already exists")
	}
	return s.repo.Create(exerciseCategory)
}

func (s *exerciseCategoryService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
