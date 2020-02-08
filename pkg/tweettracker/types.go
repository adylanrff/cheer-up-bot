package tweettracker

import (
	"encoding/json"
	"log"
)

type TwitterHandler interface {
	HandleMention() error
}

// TwitterConfig : contains the twitter API Config
type TwitterConfig struct {
	APIKey            string
	APISecretKey      string
	AccessToken       string
	AccessTokenSecret string
	Username          string
}

// Tweet : depicts the response body from the streaming Tweet API
type Tweet struct {
	ID              string `json:"id"`
	CreatedAt       string `json:"created_at"`
	Text            string `json:"text"`
	AuthorID        string `json:"author_id"`
	InReplyToUserID string `json:"in_reply_to_user_id"`
}

func ParseTweet(jsonString []byte) (*Tweet, error) {
	tweet := new(Tweet)
	err := json.Unmarshal(jsonString, &tweet)
	if err != nil {
		log.Fatal("Failed parsing tweet", err.Error())
		return nil, err
	}
	return tweet, err
}
