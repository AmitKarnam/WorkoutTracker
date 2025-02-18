package database

import (
	"gorm.io/gorm"
)

// DB interface that every database type needs to implement
type Database interface {
	GetConnection() (*gorm.DB, error)
}
