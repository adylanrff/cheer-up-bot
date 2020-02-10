package tweettracker

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/sling"
)

type TwitterAPI struct {
	config      *TwitterConfig
	handler     TwitterHandler
	bearerToken string
	tweetChan   chan string
}

func NewTwitterAPI(config *TwitterConfig, handler TwitterHandler) (*TwitterAPI, error) {
	bearerToken, err := getBearerToken(config.APIKey, config.APISecretKey)
	if err != nil {
		log.Fatal("Error getting bearer token: ", err.Error())
		return nil, err
	}
	tweetChan := make(chan string, 10)
	return &TwitterAPI{config, handler, bearerToken, tweetChan}, nil
}

func (t *TwitterAPI) doOAuth2Request(method, url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+t.bearerToken)
	req.Header.Add("User-Agent", "CheerMeUpPlease")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	return resp, nil
}

func (t *TwitterAPI) doOAuth1Request(method, url string, payload interface{}) (*http.Response, error) {
	config := oauth1.NewConfig(t.config.APIKey, t.config.APISecretKey)
	token := oauth1.NewToken(t.config.AccessToken, t.config.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	req, err := sling.New().Post(url).QueryStruct(payload).Request()
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}

	resp, err := httpClient.Do(req)
	return resp, nil
}

func (t *TwitterAPI) streamTweet() (*http.Response, error) {
	resp, err := t.doOAuth2Request(SampledStreamURL["method"], SampledStreamURL["url"], nil)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	return resp, nil
}

func (t *TwitterAPI) stream() error {
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
			t.tweetChan <- tweet.Text
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

func (t *TwitterAPI) Tweet(text string) error {
	payload := &PostTweetRequest{Status: text}
	resp, err := t.doOAuth1Request(PostTweetURL["method"], PostTweetURL["url"], payload)
	if err != nil {
		log.Fatal("Doing request: ", err)
	}
	if resp.StatusCode == 200 {
		return nil
	}
	msg, _ := ioutil.ReadAll(resp.Body)
	log.Fatal(string(msg))
	return err
}

func (t *TwitterAPI) Run() {
	var wg sync.WaitGroup
	wg.Add(2)

	// goroutine for streaming tweet data
	go func() {
		defer wg.Done()
		err := t.stream()
		if err != nil {
			log.Println("Error streaming tweets: ", err)
		}
	}()

	// gorouting for handling tweets
	go func() {
		defer wg.Done()
		for {
			err := t.handleTweet()
			if err != nil {
				log.Println("Error handling tweet: ", err)
			}
		}
	}()

	wg.Wait()
}
