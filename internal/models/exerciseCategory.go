package models

import (
	"gorm.io/gorm"
)

type ExerciseCategory struct {
	gorm.Model
	Category string `json:"category"`
}
