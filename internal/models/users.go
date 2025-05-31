package models

import (
	"errors"

	"gorm.io/gorm"
)

type UserRole string

const (
	Customer UserRole = "Customer"
	Admin    UserRole = "Admin"
)

// ValidRoles contains the allowed values for userRole
var ValidRoles = []UserRole{Customer, Admin}

// IsValid checks if the role is valid
func (r UserRole) IsValid() bool {
	for _, validRole := range ValidRoles {
		if r == validRole {
			return true
		}
	}
	return false
}

type User struct {
	gorm.Model
	Email string `gorm:"unique;not null" json:"email"`
	Name  string `json:"name"`
	Role  UserRole
}

// BeforeSave GORM hook to validate the Role field
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !u.Role.IsValid() {
		return errors.New("invalid role: must be 'Customer' or 'Admin'")
	}
	return nil
}
