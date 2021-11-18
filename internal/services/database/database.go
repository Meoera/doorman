package database

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/meoera/doorman/pkg/models"
)


type Database interface {
	Connect(credentials ...interface{}) error
	Close() error
	UserByID(id int) (*models.DatabaseUser, error)
	UserByName(name string) (*models.DatabaseUser, error)
}


type CacheDatabase interface {
	Connect(db Database, credentials ...interface{}) error
	Close() error
	UserByID(id int) (*models.DatabaseUser, error)
	UserByName(name string) (*models.DatabaseUser, error)
	AddRefreshToken(token jwt.Token) (error)
}