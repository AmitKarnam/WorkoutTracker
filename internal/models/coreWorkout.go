package models

import (
	"gorm.io/gorm"
)

type CoreWorkout struct {
	gorm.Model
	CoreID   uint         `json:"core_id"`
	Core     CoreExercise `gorm:"foreignKey:CoreID" onUpdate:"CASCADE" onDelete:"CASCADE"`
	Duration uint         `json:"duration"`
}
