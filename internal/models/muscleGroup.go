// models/muscle_group.go
package models

import (
	"gorm.io/gorm"
)

type MuscleGroup struct {
	gorm.Model
	MuscleGroup string `json:"muscle_group"`
	Description string `json:"description"`
}
