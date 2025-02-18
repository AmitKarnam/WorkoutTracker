package models

import (
	"gorm.io/gorm"
)

type CoreExercise struct {
	gorm.Model
	Name             string           `json:"name"`
	Description      string           `json:"descripton"`
	CategoryID       uint             `json:"category_id"`
	ExerciseCategory ExerciseCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
