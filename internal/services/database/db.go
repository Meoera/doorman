package database

import "github.com/meoera/doorman/internal/models"

type Database interface {
	Connect(credentials ...interface{}) error
	UserByID(id int) (*models.DatabaseUser, error)
}
