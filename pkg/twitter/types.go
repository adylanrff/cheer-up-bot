package twitter

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
	Status                    string   `url:"status"`
	InReplyToStatusID         string   `url:"in_reply_to_status_id"`
	AutoPopulateReplyMetadata bool     `url:"auto_populate_reply_metadata"`
	MediaID                   []string `url:"media_ids"`
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

type UploadMediaRequest struct {
	MediaData     string `json:"media_data"`
	MediaCategory string `json:"media_category"`
}
