package services

import (
	"context"
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type MuscleGroupService interface {
	GetAll(ctx context.Context) (*[]models.MuscleGroup, error)
	GetByID(ctx context.Context, id uint) (*models.MuscleGroup, error)
	Create(ctx context.Context, muscleGroup *models.MuscleGroup) error
	Update(ctx context.Context, id uint, input models.MuscleGroup) error
	Delete(ctx context.Context, id uint) error
}

type muscleGroupService struct {
	repo repository.MuscleGroupRepository
}

func NewMuscleGroupService(repo repository.MuscleGroupRepository) MuscleGroupService {
	return &muscleGroupService{repo: repo}
}

func (s *muscleGroupService) GetAll(ctx context.Context) (*[]models.MuscleGroup, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		return s.repo.FindAll(ctx)
	}
}

func (s *muscleGroupService) GetByID(ctx context.Context, id uint) (*models.MuscleGroup, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		return s.repo.FindByID(ctx, id)
	}
}

func (s *muscleGroupService) Create(ctx context.Context, muscleGroup *models.MuscleGroup) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		existing, err := s.repo.FindByName(ctx, muscleGroup.MuscleGroup)
		if err == nil && existing != nil {
			return fmt.Errorf("muscle group already exists")
		}
		return s.repo.Create(ctx, muscleGroup)
	}
}

func (s *muscleGroupService) Update(ctx context.Context, id uint, input models.MuscleGroup) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		muscleGroup, err := s.repo.FindByID(ctx, id)
		if err != nil {
			return err
		}

		muscleGroup.MuscleGroup = input.MuscleGroup
		muscleGroup.Description = input.Description

		err = s.repo.Update(ctx, muscleGroup)
		return err
	}
}

func (s *muscleGroupService) Delete(ctx context.Context, id uint) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request canceled: %v", ctx.Err())
	default:
		_, err := s.repo.FindByID(ctx, id)
		if err != nil {
			return err
		}
		return s.repo.Delete(ctx, id)
	}
}
