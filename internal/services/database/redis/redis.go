package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/pkg/models"
)

var (
	ErrUnspecifiedConfig       error = errors.New("you didn't specify a config")
	ErrUnspecifiedMainDatabase error = errors.New("you didn't specify a main database")
	ErrInvalidConfig           error = errors.New("the config you specified is invalid")
	ErrInvalidExpiration       error = errors.New("invalid expiration time")
)

var (
	UserStorePattern = "u:%d"
	TokenStorePattern = "rtoken:%d"
)

//the redis "connector" for the auth backend
type Redis struct {

	client redis.Client
	db     database.Database

	cfg *config.Redis
}

func (db *Redis) Client() redis.Client {
	return db.client
}

func (db *Redis) Connect(maindb database.Database, credentials ...interface{}) error {
	if len(credentials) < 1 {
		return ErrUnspecifiedConfig
	}

	if maindb != nil {
		db.db = maindb
	} else {
		return ErrUnspecifiedMainDatabase
	}

	cfg, ok := credentials[0].(*config.Redis)
	if !ok {
		return ErrInvalidConfig
	}

	db.cfg = cfg

	newClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + fmt.Sprint(cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       int(cfg.Database),
	})

	db.client = *newClient
	return nil
}

func (db *Redis) Close() error {
	return db.client.Close()
}

func (db *Redis) UserByID(id uint) (model *models.DatabaseUser, err error) {
	exists := db.client.Exists(context.Background(), fmt.Sprintf(UserStorePattern, id))
	if exists.Err() != nil {
		return nil, exists.Err()
	} else if exists.Val() != 1 {
		model, err := db.db.UserByID(id)
		if err != nil {
			return model, err
		} else if model == nil {
			return model, err
		}

		tNow := time.Now()
		db.client.SetEX(
			context.Background(),
			fmt.Sprintf(UserStorePattern, id),
			model,
			time.Duration(tNow.Unix()-tNow.Add(time.Duration(db.cfg.StandartExpiration)).Unix())*time.Second,
		)
	}

	md := db.client.Get(context.Background(), fmt.Sprintf(UserStorePattern, id))
	fmt.Println(md.Val())

	return
}

func (db *Redis) UserByName(name string) (*models.DatabaseUser, error) {
	return db.db.UserByName(name)
}


func (db *Redis) AddRefreshToken(token string, uid, exp uint) (err error) {
	_, err = db.client.SetEX(context.Background(), fmt.Sprintf(TokenStorePattern, uid), token, time.Duration(int64(exp)-time.Now().Unix()) * time.Second).Result()
	if err != nil {
		return
	}


	return
}
