package main

import (
	"flag"
	"fmt"

	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/inits"
	"github.com/meoera/doorman/pkg/hasher"
)

const DEFAULT_CONFIG_PATH string = "./config.json"

var configFlag = flag.String("c", DEFAULT_CONFIG_PATH, "Config File Path")

var devFlag = flag.Bool("d", false, "Developing mode")

func main() {
	flag.Parse()

	hasher.HashPassword("test", nil)

	cfg, err := config.Parse(*configFlag)
	if err != nil {
		fmt.Println("An error occured while loading your config file:", err)
	}
	fmt.Println("Config successfully loaded!")
	
	db, err := inits.Database(cfg.Database.MySQL)
	if err != nil {
		panic(err)
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
