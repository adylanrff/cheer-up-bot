package tweettracker

import (
	"encoding/json"
	"io"
)

// TweetResponse : depicts the response body from the streaming Tweet API
type Tweet struct {
	ID              string `json:"id"`
	CreatedAt       string `json:"created_at"`
	Text            string `json:"text"`
	AuthorID        string `json:"author_id"`
	InReplyToUserID string `json:"in_reply_to_user_id"`
}

type TweetResponse struct {
	Data Tweet `json:"data"`
}

func ParseTweet(reader io.Reader) (Tweet, error) {
	tweet := new(TweetResponse)
	json.NewDecoder(reader).Decode(&tweet)
	return tweet.Data, nil
}
