package twitter

// FilteredStreamURL : URL for Filtered Twitter Streaming API
var FilteredStreamURL = map[string]string{
	"method": "GET",
	"url":    "https://api.twitter.com/labs/1/tweets/stream/filter",
}

// SampledStreamURL : URL for Filtered Twitter Streaming API
var SampledStreamURL = map[string]string{
	"method": "GET",
	"url":    "https://api.twitter.com/labs/1/tweets/stream/sample",
}

// RulesURL : URL for Twitter Rule for Streaming API
var RulesURL = map[string]string{
	"method": "POST",
	"url":    "https://api.twitter.com/labs/1/tweets/stream/filter/rules",
}

// BearerTokenURL : URL for OAuth2 Authorization
var BearerTokenURL = map[string]string{
	"method": "POST",
	"url":    "https://api.twitter.com/oauth2/token",
}

var PostTweetURL = map[string]string{
	"method": "POST",
	"url":    "https://api.twitter.com/1.1/statuses/update.json",
}

var AddStreamFilterRuleURL = map[string]string{
	"method": "POST",
	"url":    "https://api.twitter.com/labs/1/tweets/stream/filter/rules",
}

var GetStreamFilterRuleURL = map[string]string{
	"method": "GET",
	"url":    "https://api.twitter.com/labs/1/tweets/stream/filter/rules",
}

var GetUserInfoURL = map[string]string{
	"method": "GET",
	"url":    "https://api.twitter.com/1.1/users/lookup.json",
}

var UploadImageURL = map[string]string{
	"method": "POST",
	"url":    "https://upload.twitter.com/1.1/media/upload.json",
}
