package config

import "gopkg.in/ini.v1"

// Config : the config of the project
type Config struct {
	APIKey            string
	APISecretKey      string
	AccessToken       string
	AccessTokenSecret string
}

func NewConfig(cfg *ini.File) *Config {

	config := Config{
		cfg.Section("twitter").Key("api_key").String(),
		cfg.Section("twitter").Key("api_secret_key").String(),
		cfg.Section("twitter").Key("access_token").String(),
		cfg.Section("twitter").Key("access_token_secret").String(),
	}
	return &config
}
