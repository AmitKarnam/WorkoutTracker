package models

import (
	"gorm.io/gorm"
)

type CoreExercise struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"descripton"`
}
