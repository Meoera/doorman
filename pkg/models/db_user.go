package models

import "gorm.io/gorm"

// The user model used by the database
type DatabaseUser struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	PasswordHash string
	Salt string `gorm:"<-:create"`
}