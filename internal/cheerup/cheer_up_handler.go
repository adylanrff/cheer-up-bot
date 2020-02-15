package cheerup

import (
	"fmt"

	"github.com/adylanrff/cheer-up-bot/pkg/tweettracker"
)

type CheerUpHandler struct {
}

func NewCheerUpHandler() *CheerUpHandler {
	return &CheerUpHandler{}
}

func (c *CheerUpHandler) HandleMention(twitterAPI *tweettracker.TwitterAPI, text string) error {
	fmt.Println(text)
	return nil
}
