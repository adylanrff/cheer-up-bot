package tweettracker

type TwitterHandler interface {
	HandleMention(*TwitterAPI, *Tweet) error
}

// TwitterConfig : contains the twitter API Config
type TwitterConfig struct {
	APIKey            string
	APISecretKey      string
	AccessToken       string
	AccessTokenSecret string
	Username          string
}

type BearerTokenRequest struct {
	GrantType string `url:"grant_type"`
}

type BearerTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type PostTweetRequest struct {
	Status                    string `url:"status"`
	InReplyToStatusID         string `url:"in_reply_to_status_id"`
	AutoPopulateReplyMetadata bool   `json:"auto_populate_reply_metadata"`
}

type TweetFilterRule struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

type TweetFilterRuleRequest struct {
	Add []TweetFilterRule `json:"add"`
}

type TweetFilterRuleResponse struct {
	Data []TweetFilterRule `json:"data"`
}

type GetUserRequest struct {
	ScreenName []string `url:"screen_name"`
	UserID     []string `url:"user_id"`
}
