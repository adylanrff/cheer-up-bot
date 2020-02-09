package tweettracker

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type TwitterAPI struct {
	config      *TwitterConfig
	handler     TwitterHandler
	bearerToken string
}

func NewTwitterAPI(config *TwitterConfig, handler TwitterHandler) (*TwitterAPI, error) {
	bearerToken, err := getBearerToken(config.APIKey, config.APISecretKey)
	if err != nil {
		log.Fatal("Error getting bearer token: ", err.Error())
		return nil, err
	}
	return &TwitterAPI{config, handler, bearerToken}, nil
}

func (t *TwitterAPI) streamTweet() (*http.Response, error) {
	req, err := http.NewRequest("GET", SampledStreamURL, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+t.bearerToken)
	req.Header.Add("User-Agent", "CheerMeUpPlease")
	resp, err := http.DefaultClient.Do(req)
	return resp, nil
}

func (t *TwitterAPI) stream() error {
	resp, err := t.streamTweet()
	if err != nil {
		log.Fatal("Error streaming tweet: ", err.Error())
		return err
	}
	fmt.Println(resp.StatusCode)
	for {
		tweet, err := ParseTweet(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(tweet.Text)
	}

	return nil
}

func (t *TwitterAPI) Run() error {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		t.stream()
	}()

	wg.Wait()
	return nil
}
