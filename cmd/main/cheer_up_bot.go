package main

import (
	"fmt"
	"os"

	"github.com/adylanrff/cheer-up-bot/internal/config"
	"gopkg.in/ini.v1"
)

// CONFIGFILEPATH is a file path containing the config of the project
const CONFIGFILEPATH string = "config/config.ini"

func main() {
	cfg, err := ini.Load(CONFIGFILEPATH)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println(config.NewConfig(cfg))
	fmt.Println("Running bot...")
}
