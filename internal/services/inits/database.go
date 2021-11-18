package inits

import (
	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/internal/services/database/mysql"
	"github.com/meoera/doorman/internal/services/database/redis"
)

func Database(cfg *config.MySQL) (db *mysql.MySQL, err error) {
	db = &mysql.MySQL{}

	err = db.Connect(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	return
}

func CacheDatabase(cfg *config.Redis, middleware database.Database) (db *redis.Redis, err error) {
	db = &redis.Redis{}

	err = db.Connect(middleware, cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)

	return
}