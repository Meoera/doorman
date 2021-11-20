package database

import (
	"github.com/meoera/doorman/pkg/models"
)


type Database interface {
	Connect(credentials ...interface{}) error
	Close() error
	UserByID(id uint) (*models.DatabaseUser, error)
	UserByName(name string) (*models.DatabaseUser, error)
}


type CacheDatabase interface {
	Connect(db Database, credentials ...interface{}) error
	Close() error
	UserByID(id uint) (*models.DatabaseUser, error)
	UserByName(name string) (*models.DatabaseUser, error)
	AddRefreshToken(token string, uid, exp uint) (error)
}