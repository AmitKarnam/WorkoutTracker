package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	UserID     uint `gorm:"uniqueIndex"`
	User       User `gorm:"foreignKey:UserID" onUpdate:"CASCADE" onDelete:"CASCADE`
	Token      string
	ExpiryTime time.Time
}
