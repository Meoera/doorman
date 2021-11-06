package mysql

import (
	"errors"
	"fmt"

	"github.com/meoera/doorman/internal/models"
	"github.com/meoera/doorman/internal/services/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MYSQL_DSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

var (
	ErrInvalidHost     = errors.New("The specified MySQL host is invalid")
	ErrInvalidPort     = errors.New("The specified MySQL port is invalid")
	ErrInvalidUser     = errors.New("The specified MySQL user is invalid!")
	ErrInvalidPassword = errors.New("The specified MySQL password is invalid!")
	ErrInvalidDbName   = errors.New("The specified MySQL database-name is invalid!")
)

type MySQL struct {
	database.Database

	connector *gorm.DB
}

func (db *MySQL) Connect(credentials ...interface{}) (err error) {
	host, ok := credentials[0].(string)
	if !ok {
		return ErrInvalidHost
	}

	port, ok := credentials[1].(uint)
	if !ok {
		return ErrInvalidPort
	}

	user, ok := credentials[2].(string)
	if !ok {
		return ErrInvalidUser
	}

	password, ok := credentials[3].(string)
	if !ok {
		return ErrInvalidPassword
	}

	databaseName, ok := credentials[4].(string)
	if !ok {
		return ErrInvalidDbName
	}

	dsn_formatted := fmt.Sprintf(MYSQL_DSN, user, password, host, port, databaseName)

	conn, err := gorm.Open(mysql.Open(dsn_formatted), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.connector.AutoMigrate(&models.DatabaseUser{})
	if err != nil {
		return err
	}

	db.connector = conn
	return
}

func (db *MySQL) UserByID(id int) (result *models.DatabaseUser, err error) {
	err = db.connector.First(result, id).Error
	return
}
