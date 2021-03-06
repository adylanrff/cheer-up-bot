package cheerup

import (
	"fmt"

	"github.com/adylanrff/cheer-up-bot/internal/config"
	"github.com/adylanrff/cheer-up-bot/pkg/twitter"
)

func NewCheerUpRules(cfg *config.Config) []twitter.TweetFilterRule {
	var rules []twitter.TweetFilterRule

	mentionRule := twitter.TweetFilterRule{
		Value: fmt.Sprintf("@%s semangat", cfg.TwitterUsername),
		Tag:   "semangat",
	}

	rules = append(rules, mentionRule)
	return rules
}
