package twitter

import (
	"bytes"
	"encoding/json"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// TweetResponse : depicts the response body from the streaming Tweet API
type Tweet struct {
	ID                        string   `json:"id"`
	CreatedAt                 string   `json:"created_at"`
	Text                      string   `json:"text"`
	AuthorID                  string   `json:"author_id"`
	InReplyToStatusID         string   `json:"in_reply_to_status_id"`
	AutoPopulateReplyMetadata bool     `json:"auto_populate_reply_metadata"`
	MediaID                   []string `json:"media_ids"`
	Image                     []image.Image
}

type TweetResponse struct {
	Data Tweet `json:"data"`
}

func ParseTweet(reader io.Reader) (Tweet, error) {
	tweet := new(TweetResponse)
	json.NewDecoder(reader).Decode(&tweet)
	return tweet.Data, nil
}

func (t *TwitterAPI) AddRule() error {
	request := TweetFilterRuleRequest{Add: t.rules}
	payload, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Error marshalling Twitter Rule: ", err)
		return err
	}

	resp, err := t.doOAuth2Request(AddStreamFilterRuleURL["method"], AddStreamFilterRuleURL["url"], bytes.NewBuffer(payload))
	if resp.StatusCode != 201 {
		log.Fatal("Adding rule failed: ", resp.StatusCode)
	}

	return nil
}

func (t *TwitterAPI) GetRules() ([]TweetFilterRule, error) {
	var rules TweetFilterRuleResponse

	resp, err := t.doOAuth2Request(GetStreamFilterRuleURL["method"], GetStreamFilterRuleURL["url"], nil)
	if err != nil {
		log.Fatal("Failed getting rules: ", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
		return nil, err
	}

	err = json.Unmarshal(body, &rules)
	if err != nil {
		log.Fatal("Error unmarshalling rules JSON: ", err)
		return nil, err
	}

	return rules.Data, nil
}

func (t *TwitterAPI) streamTweet() (*http.Response, error) {
	resp, err := t.doOAuth2Request(FilteredStreamURL["method"], FilteredStreamURL["url"], nil)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	return resp, nil
}

func (t *TwitterAPI) stream() error {
	err := t.AddRule()
	if err != nil {
		log.Fatal("Error adding rule: ", err)
	}

	resp, err := t.streamTweet()
	if err != nil {
		log.Fatal("Error streaming tweet: ", err.Error())
		return err
	}
	for {
		tweet, err := ParseTweet(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if len(tweet.Text) > 0 {
			t.tweetChan <- &tweet
		}
	}
}

func (t *TwitterAPI) handleTweet() error {
	tweet := <-t.tweetChan
	err := t.handler.HandleMention(t, tweet)
	if err != nil {
		return err
	}
	return nil
}

func (t *TwitterAPI) Tweet(tweet *Tweet) error {
	if len(tweet.Image) > 0 {
		for _, img := range tweet.Image {
			media, err := t.UploadImage(img)
			log.Fatal("Error uploading image: ", err)
			tweet.MediaID = append(tweet.MediaID, media.MediaID)
		}
	}

	payload := &PostTweetRequest{
		Status:                    tweet.Text,
		InReplyToStatusID:         tweet.InReplyToStatusID,
		AutoPopulateReplyMetadata: true,
		MediaID:                   tweet.MediaID,
	}

	resp, err := t.doOAuth1Request(PostTweetURL["method"], PostTweetURL["url"], payload)
	if err != nil {
		log.Fatal("Doing request: ", err)
	}
	if resp.StatusCode != 200 {
		msg, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(msg))
		return err
	}
	log.Println("Tweet: ", payload.Status)
	return nil
}
