// models/workout.go
package models

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model          // Embedding gorm.Model
	ExerciseID uint     `json:"exercise_id"` // Foreign key for Exercise
	Reps       int64    `json:"reps"`
	Weight     int64    `json:"weight"`
	Exercise   Exercise `gorm:"foreignKey:ExerciseID","constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association
}
