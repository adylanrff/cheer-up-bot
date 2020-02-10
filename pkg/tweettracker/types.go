package tweettracker

type TwitterHandler interface {
	HandleMention(*TwitterAPI, string) error
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
	Status            string `url:"status"`
	InReplyToStatusID string `url:"in_reply_to_status_id"`
}
