package tweettracker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TwitterAPI struct {
	config      *TwitterConfig
	handler     TwitterHandler
	bearerToken string
}

func NewTwitterAPI(config *TwitterConfig, handler TwitterHandler) *TwitterAPI {
	bearerToken, err := getBearerToken(config.APIKey, config.APISecretKey)
	if err != nil {
		log.Fatal("Error getting bearer token: ", err.Error())
	}
	fmt.Println(bearerToken)
	return &TwitterAPI{config, handler, bearerToken}
}

func (t *TwitterAPI) streamTweet() (*http.Response, error) {
	req, err := http.NewRequest("GET", StreamURL, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	return resp, nil
}

func (t *TwitterAPI) Connect() error {
	resp, err := t.streamTweet()
	if err != nil {
		log.Fatal("Error streaming tweet: ", err.Error())
		return err
	}
	for {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		if resp.StatusCode == 200 {
			fmt.Println(ParseTweet(body))
		}
	}

	return nil
}
