package services

import (
	"context"
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
)

type StrengthExerciseService interface {
	GetAll(ctx context.Context) (*[]models.StrengthExercise, error)
	GetByMuscleGroup(ctx context.Context, mucleGroupName string) (*[]models.StrengthExercise, error)
	GetByID(ctx context.Context, id uint) (*models.StrengthExercise, error)
	Create(ctx context.Context, strengthExercise *models.StrengthExercise) error
	Update(ctx context.Context, id uint, input models.StrengthExercise) error
	Delete(ctx context.Context, id uint) error
}

type strengthExerciseService struct {
	repo repository.StrengthExerciseRepository
}

func NewStrengthExerciseService(repo repository.StrengthExerciseRepository) StrengthExerciseService {
	return &strengthExerciseService{repo: repo}
}

func (s *strengthExerciseService) GetAll(ctx context.Context) (*[]models.StrengthExercise, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		return s.repo.FindAll(ctx)
	}
}

func (s *strengthExerciseService) GetByMuscleGroup(ctx context.Context, muscleGroupName string) (*[]models.StrengthExercise, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		return s.repo.FindByMuscleGroup(ctx, muscleGroupName)
	}
}

func (s *strengthExerciseService) GetByID(ctx context.Context, id uint) (*models.StrengthExercise, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		return s.repo.FindByID(ctx, id)
	}
}

func (s *strengthExerciseService) Create(ctx context.Context, strengthExercise *models.StrengthExercise) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		return s.repo.Create(ctx, strengthExercise)
	}
}

func (s *strengthExerciseService) Update(ctx context.Context, id uint, input models.StrengthExercise) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		strengthExercise, err := s.repo.FindByID(ctx, id)
		if err != nil {
			return err
		}

		strengthExercise.Name = input.Name
		strengthExercise.Description = input.Description
		strengthExercise.BodyWeight = input.BodyWeight
		strengthExercise.MuscleGroupID = input.ID

		err = s.repo.Update(ctx, strengthExercise)
		return err
	}
}

func (s *strengthExerciseService) Delete(ctx context.Context, id uint) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("request cancelled: %v", ctx.Err())
	default:
		_, err := s.repo.FindByID(ctx, id)
		if err != nil {
			return err
		}
		return s.repo.Delete(ctx, id)
	}
}
