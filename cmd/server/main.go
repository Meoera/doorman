package main

import (
	"flag"
	"fmt"

	"github.com/meoera/doorman/internal/services/config"
	"github.com/meoera/doorman/internal/services/inits"
)

const DEFAULT_CONFIG_PATH string = "./config.json"

var configFlag = flag.String("c", DEFAULT_CONFIG_PATH, "Config File Path")

var devFlag = flag.Bool("d", false, "Developing mode")

func main() {
	flag.Parse()

	cfg, err := config.Parse(*configFlag)
	if err != nil {
		fmt.Println("An error occured while loading your config file:", err)
	}
	fmt.Println("Config successfully loaded!")

	err = inits.InitializeWeb(cfg.Web, cfg.Database.Redis, *devFlag)
	if err != nil {
		panic(err)
	}
}
