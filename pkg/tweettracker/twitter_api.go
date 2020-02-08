package tweettracker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TwitterAPI struct {
	config     *TwitterConfig
	handler    TwitterHandler
	httpClient *http.Client
}

func NewTwitterAPI(config *TwitterConfig, handler TwitterHandler) *TwitterAPI {
	client := http.Client{}
	return &TwitterAPI{config, handler, &client}
}

func getBearerToken() string {
	// TODO: Implement This
	return "BEARER_TOKEN"
}

func (t *TwitterAPI) streamTweet() (*http.Response, error) {
	req, err := http.NewRequest("GET", StreamURL, nil)
	req.BasicAuth()
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}
	resp, err := t.httpClient.Do(req)
	return resp, nil
}

func (t *TwitterAPI) Connect() error {
	// TODO: Add Bearer Token Auth
	bearerToken := getBearerToken()
	fmt.Println(bearerToken)

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
