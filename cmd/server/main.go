package main

import (
	"flag"
	"fmt"

	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/database"
	"github.com/meoera/doorman/internal/services/database/mock"
	"github.com/meoera/doorman/internal/services/inits"
	"github.com/meoera/doorman/pkg/hasher"
	"github.com/meoera/doorman/pkg/models"
)

const DEFAULT_CONFIG_PATH string = "./config.json"

var configFlag = flag.String("c", DEFAULT_CONFIG_PATH, "Config File Path")

var devFlag = flag.Bool("d", false, "Developing mode")
var useMysql = flag.Bool("mysql", false, "Use mysql")
var useMock = flag.Bool("mock", false, "Use mock db")

func main() {
	flag.Parse()

	if *useMock && !*devFlag {
		panic("mock only usable in dev mode")
	}

	if !*useMock && !*useMysql {
		panic("you have to use a db")
	}

	cfg, err := config.Parse(*configFlag)
	if err != nil {
		fmt.Println("An error occured while loading your config file:", err)
	}

	fmt.Println("Config successfully loaded!")

	var db database.Database

	if *useMysql {
		db, err = inits.MySql(cfg.Database.MySQL)
		if err != nil {
			panic(err)
		}
	} else if *useMock {
		db = &mock.MockDatabase{}
		testpw, testsalt, err := hasher.HashPassword("testpassword", nil)
		if err != nil {
			panic("creating the passwordhash for your mock user failed")
		}
		db.Connect([]models.DatabaseUser{
			{
				Username: "testuser",
				Email: "test@email.com",
				PasswordHash: testpw,
				Salt: testsalt,
			},
		})
	}

	defer db.Close()

	cdb, err := inits.CacheDatabase(cfg.Database.Redis, db)
	if err != nil {
		panic(err)
	}
	defer cdb.Close()

	err = inits.Web(cfg.Web, db, cdb, *devFlag)
	if err != nil {
		panic(err)
	}
}
