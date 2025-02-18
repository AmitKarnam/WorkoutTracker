package models

import (
	"gorm.io/gorm"
)

type StrengthWorkout struct {
	gorm.Model
	ExerciseID  uint             `json:"exercise_id"`
	Exercise    StrengthExercise `gorm:"foreignKey:ExerciseID" onUpdate:"CASCADE" onDelete:"CASCADE"`
	Weight      float64          `json:"weight"`
	Repetetions uint             `json:"repetetions"`
}
