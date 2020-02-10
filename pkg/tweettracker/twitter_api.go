package tweettracker

import (
	"log"
	"net/http"
	"sync"
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

func (t *TwitterAPI) doRequest(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+t.bearerToken)
	req.Header.Add("User-Agent", "CheerMeUpPlease")
	resp, err := http.DefaultClient.Do(req)
	return resp, nil
}

func (t *TwitterAPI) streamTweet() (*http.Response, error) {
	resp, err := t.doRequest(SampledStreamURL["method"], SampledStreamURL["url"])
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

func (t *TwitterAPI) Tweet() error {
	return nil
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
