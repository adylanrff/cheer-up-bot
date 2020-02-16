package cheerup

import (
	"fmt"
	"log"
	s "strings"

	"github.com/adylanrff/cheer-up-bot/pkg/twitter"
)

type CheerUpHandler struct {
}

func NewCheerUpHandler() *CheerUpHandler {
	return &CheerUpHandler{}
}

func (c *CheerUpHandler) HandleMention(twitterAPI *twitter.TwitterAPI, tweet *twitter.Tweet) error {
	var response *twitter.Tweet
	if s.Contains(s.ToLower(tweet.Text), "semangat") {
		response = c.CreateSemangatResponse(twitterAPI, tweet)
	}

	err := twitterAPI.Tweet(response)
	if err != nil {
		log.Fatal("Error responding to mentions: ", err)
	}

	return nil
}

func (c *CheerUpHandler) CreateSemangatResponse(twitterAPI *twitter.TwitterAPI, tweet *twitter.Tweet) *twitter.Tweet {
	user, err := twitterAPI.GetUser(tweet.AuthorID)
	if err != nil {
		log.Println("Error fetching user: ", err)
	}
	response := &twitter.Tweet{
		Text:              fmt.Sprintf("Semangat %s!", user.Name),
		InReplyToStatusID: tweet.ID,
	}

	return response
}
