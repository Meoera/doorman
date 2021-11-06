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
}

type Database struct {
	Redis *Redis `json:"redis"`
}

type Redis struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	Username string `json:"username"`
	Password string `json:"password"`
	Database int    `json:"database"`
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
			Redis: &Redis{},
		},
	}

	err = json.Unmarshal(data, cfg)
	return
}
