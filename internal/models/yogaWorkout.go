package models

import (
	"gorm.io/gorm"
)

type YogaWorkout struct {
	gorm.Model
	YogaID   uint         `json:"yoga_id"`
	Yoga     YogaExercise `gorm:"foreignKey:YogaID" onUpdate:"CASCADE" onDelete:"CASCADE"`
	Duration uint         `json:"duration"`
}
