package services

import (
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type MuscleGroupService interface {
	GetAll() ([]models.MuscleGroup, error)
	GetByID(id uint) (*models.MuscleGroup, error)
	Create(muscleGroup *models.MuscleGroup) error
	Update(id uint, input models.MuscleGroup) (*models.MuscleGroup, error)
	Delete(id uint) error
}

type muscleGroupService struct {
	repo repository.MuscleGroupRepository
}

func NewMuscleGroupService(repo repository.MuscleGroupRepository) MuscleGroupService {
	return &muscleGroupService{repo: repo}
}

func (s *muscleGroupService) GetAll() ([]models.MuscleGroup, error) {
	return s.repo.FindAll()
}

func (s *muscleGroupService) GetByID(id uint) (*models.MuscleGroup, error) {
	return s.repo.FindByID(id)
}

func (s *muscleGroupService) Create(muscleGroup *models.MuscleGroup) error {
	existing, err := s.repo.FindByName(muscleGroup.MuscleGroup)
	if err == nil && existing != nil {
		return fmt.Errorf("muscle group already exists")
	}
	return s.repo.Create(muscleGroup)
}

func (s *muscleGroupService) Update(id uint, input models.MuscleGroup) (*models.MuscleGroup, error) {
	muscleGroup, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	muscleGroup.MuscleGroup = input.MuscleGroup
	muscleGroup.Description = input.Description

	err = s.repo.Update(muscleGroup)
	return muscleGroup, err
}

func (s *muscleGroupService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
