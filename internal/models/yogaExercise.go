package models

import (
	"gorm.io/gorm"
)

type YogaExercise struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
