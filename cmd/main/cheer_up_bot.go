package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adylanrff/cheer-up-bot/internal/cheerup"
	"github.com/adylanrff/cheer-up-bot/internal/config"
	"github.com/adylanrff/cheer-up-bot/pkg/tweettracker"
	"gopkg.in/ini.v1"
)

// CONFIGFILEPATH is a file path containing the config of the project
const CONFIGFILEPATH string = "config/config.ini"

func main() {
	cfgFile, err := ini.Load(CONFIGFILEPATH)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	appConfig := config.NewConfig(cfgFile)

	twitterConfig := tweettracker.TwitterConfig{
		appConfig.APIKey,
		appConfig.APISecretKey,
		appConfig.AccessToken,
		appConfig.AccessTokenSecret,
		appConfig.TwitterUsername,
	}

	cheerUpHandler := cheerup.NewCheerUpHandler()
	cheerUpFilterRules := cheerup.NewCheerUpRules(appConfig)

	fmt.Println("Running bot...")
	tracker, err := tweettracker.NewTwitterAPI(&twitterConfig, cheerUpHandler, cheerUpFilterRules)
	if err != nil {
		log.Panicln("Failed initiating tweet tracker")
	}
	// Validate rules
	rules, err := tracker.GetRules()
	if err != nil {
		log.Panicln("Error getting rules: ", err)
	}

	fmt.Println(rules)

	tracker.Run()
}
