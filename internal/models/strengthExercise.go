// models/exercise.go
package models

import (
	"gorm.io/gorm"
)

type StrengthExercise struct {
	gorm.Model                        // Embedding gorm.Model
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	BodyWeight       bool             `json:"body_weight"`
	MuscleGroupID    uint             `json:"muscle_group_id"`                                                                      // Foreign key for MuscleGroup
	MuscleGroup      MuscleGroup      `gorm:"foreignKey:MuscleGroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association
	CategoryID       uint             `json:"category_id"`                                                                          // Foreign key for Category
	ExerciseCategory ExerciseCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`    // Association
}
