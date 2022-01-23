package config

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrInvalidPath = errors.New("the path you specified is invalid")
)

type Config struct {
	Web      *Web      `json:"web"`
	Database *Database `json:"database"`
}

type Web struct {
	Host string `json:"host"`
	SingingSecret string `json:"secret"`
	AccessTokenExpiry uint `json:"accesstoken_expiry"`
	RefreshTokenExpiry uint `json:"refreshtoken_expiry"`
}

type Database struct {
	MySQL *MySQL `json:"mysql"`
	Redis *Redis `json:"redis"`
}

type MySQL struct {
	Host string `json:"host"`
	Port uint `json:"port"`

	User     string `json:"user"`
	Password string `json:"password"`
	Database uint `json:"database"`
}

type Redis struct {
	Host string `json:"host"`
	Port uint   `json:"port"`

	Username string `json:"username"`
	Password string `json:"password"`
	Database uint   `json:"database"`

	StandartExpiration int64 `json:"standart_expiration"`
}

func Parse(path string) (cfg *Config, err error) {
	if path == "" {
		return cfg, ErrInvalidPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}


	cfg = &Config{
		Web: &Web{},
		Database: &Database{
			MySQL: &MySQL{},
			Redis: &Redis{},
		},
	}

	err = json.Unmarshal(data, cfg)
	return
}
