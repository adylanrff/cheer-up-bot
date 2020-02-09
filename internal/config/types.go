package config

import "gopkg.in/ini.v1"

// Config : the config of the project
type Config struct {
	APIKey            string
	APISecretKey      string
	AccessToken       string
	AccessTokenSecret string
	HTTPClientTimeout int
}

func NewConfig(cfg *ini.File) *Config {

	config := Config{
		cfg.Section("twitter").Key("api_key").String(),
		cfg.Section("twitter").Key("api_secret_key").String(),
		cfg.Section("twitter").Key("access_token").String(),
		cfg.Section("twitter").Key("access_token_secret").String(),
		cfg.Section("app").Key("http_client_timeout").InInt(0, []int{5, 10}),
	}
	return &config
}
