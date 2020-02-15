package cheerup

import (
	"fmt"

	"github.com/adylanrff/cheer-up-bot/internal/config"
	"github.com/adylanrff/cheer-up-bot/pkg/tweettracker"
)

func NewCheerUpRules(cfg *config.Config) []tweettracker.TweetFilterRule {
	var rules []tweettracker.TweetFilterRule

	mentionRule := tweettracker.TweetFilterRule{
		Value: fmt.Sprintf("@%s", cfg.TwitterUsername),
		Tag:   "mention",
	}

	rules = append(rules, mentionRule)
	return rules
}
