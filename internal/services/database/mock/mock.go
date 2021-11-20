package mock

import (
	"errors"

	"github.com/meoera/doorman/pkg/models"
)


type MockDatabase struct {
	users map[uint]*models.DatabaseUser
	tokens map[uint]string
}


func (db *MockDatabase) Connect(credentials ...interface{}) error {
	if len(credentials) == 0 {
		return errors.New("you must specify mock users")
	}
	


	mockUsers, ok := credentials[0].([]*models.DatabaseUser)
	if !ok {
		return errors.New("you must specifiy valid mock users")
	}

	db.users = map[uint]*models.DatabaseUser{}

	for _, user := range mockUsers {
		db.users[user.ID] = user
	}

	return nil
}
func (db *MockDatabase) Close() error {
	return nil
}
func (db *MockDatabase) UserByID(id uint) (*models.DatabaseUser, error) {
	return db.users[id], nil
}
func (db *MockDatabase) UserByName(name string) (*models.DatabaseUser, error) {
	for _, u := range db.users {
		if u.Username == name {
			return u, nil
		}
	}
	return nil, nil
}

func (db *MockDatabase) AddRefreshToken(token string, uid, exp uint) (error){
	db.tokens[uid] = token
	
	return nil
}