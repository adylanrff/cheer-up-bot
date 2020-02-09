package tweettracker

import (
	"encoding/json"
	"log"
)

type TwitterHandler interface {
	HandleMention(*TwitterAPI, string) error
}

// TwitterConfig : contains the twitter API Config
type TwitterConfig struct {
	APIKey            string
	APISecretKey      string
	AccessToken       string
	AccessTokenSecret string
	Username          string
}

type BearerTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func ParseBearerTokenResponse(jsonString []byte) (*BearerTokenResponse, error) {
	bearerTokenResponse := new(BearerTokenResponse)
	err := json.Unmarshal(jsonString, &bearerTokenResponse)
	if err != nil {
		log.Fatal("Failed parsing tweet", err.Error())
		return nil, err
	}
	return bearerTokenResponse, nil
}
