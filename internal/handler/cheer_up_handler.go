package handler

import (
	"fmt"
	"strings"

	"github.com/adylanrff/cheer-up-bot/pkg/tweettracker"
)

type CheerUpHandler struct {
}

func NewCheerUpHandler() *CheerUpHandler {
	return &CheerUpHandler{}
}

func (c *CheerUpHandler) HandleMention(twitterAPI *tweettracker.TwitterAPI, text string) error {
	if strings.Contains(text, "Oscar") {
		fmt.Println("Tweeted")
		twitterAPI.Tweet(text)
	}
	return nil
}
