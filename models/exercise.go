// models/exercise.go
package models

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model                // Embedding gorm.Model
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	MuscleGroupID uint        `json:"muscle_group_id"`          // Foreign key for MuscleGroup
	MuscleGroup   MuscleGroup `gorm:"foreignKey:MuscleGroupID"` // Association
}
