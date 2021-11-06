package models

import "gorm.io/gorm"

type DatabaseUser struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	PasswordHash string ``
	Salt string `gorm:"<-:create"`
}