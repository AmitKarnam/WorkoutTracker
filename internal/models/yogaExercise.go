package models

import (
	"gorm.io/gorm"
)

type YogaExercise struct {
	gorm.Model
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	CategoryID   uint             `json:"category_id"`
	YogaCategory ExerciseCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
