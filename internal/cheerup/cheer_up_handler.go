package cheerup

import (
	"fmt"
	"log"
	s "strings"

	"github.com/adylanrff/cheer-up-bot/pkg/tweettracker"
)

type CheerUpHandler struct {
}

func NewCheerUpHandler() *CheerUpHandler {
	return &CheerUpHandler{}
}

func (c *CheerUpHandler) HandleMention(twitterAPI *tweettracker.TwitterAPI, tweet *tweettracker.Tweet) error {
	var response *tweettracker.Tweet
	if s.Contains(s.ToLower(tweet.Text), "semangat") {
		response = c.CreateSemangatResponse(twitterAPI, tweet)
	}

	err := twitterAPI.Tweet(response)
	if err != nil {
		log.Fatal("Error responding to mentions: ", err)
	}

	return nil
}

func (c *CheerUpHandler) CreateSemangatResponse(twitterAPI *tweettracker.TwitterAPI, tweet *tweettracker.Tweet) *tweettracker.Tweet {
	user, err := twitterAPI.GetUser(tweet.AuthorID)
	if err != nil {
		log.Println("Error fetching user: ", err)
	}
	response := &tweettracker.Tweet{
		Text:              fmt.Sprintf("Semangat %s!", user.Name),
		InReplyToStatusID: tweet.ID,
	}

	return response
}
