package cheerup

import (
	"fmt"
	"log"

	"github.com/adylanrff/cheer-up-bot/pkg/twitter"
)

type CheerUpHandler struct {
}

func NewCheerUpHandler() *CheerUpHandler {
	return &CheerUpHandler{}
}

func (c *CheerUpHandler) HandleMention(twitterAPI *twitter.TwitterAPI, tweet *twitter.Tweet) error {
	var response *twitter.Tweet
	for _, rule := range tweet.MatchingRules {
		if rule.Tag == "semangat" {
			response = c.CreateSemangatResponse(twitterAPI, tweet)
		}
	}

	if response != nil {
		err := twitterAPI.Tweet(response)
		if err != nil {
			log.Fatal("Error responding to mentions: ", err)
		}
	}

	return nil
}

func (c *CheerUpHandler) CreateSemangatResponse(twitterAPI *twitter.TwitterAPI, tweet *twitter.Tweet) *twitter.Tweet {
	userID := tweet.InReplyToUserID
	if userID == "" {
		userID = tweet.AuthorID
	}
	user, err := twitterAPI.GetUser(userID)

	if err != nil {
		log.Println("Error fetching user: ", err)
	}

	response := &twitter.Tweet{
		Text:              fmt.Sprintf("Semangat %s!", user.Name),
		InReplyToStatusID: tweet.ID,
	}

	return response
}
