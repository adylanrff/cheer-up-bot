package tweettracker

import (
	"bytes"
	"encoding/json"
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
	rules       []TweetFilterRule
}

func NewTwitterAPI(config *TwitterConfig, handler TwitterHandler, rules []TweetFilterRule) (*TwitterAPI, error) {
	bearerToken, err := getBearerToken(config.APIKey, config.APISecretKey)
	if err != nil {
		log.Fatal("Error getting bearer token: ", err.Error())
		return nil, err
	}

	tweetChan := make(chan string, 10)
	return &TwitterAPI{config, handler, bearerToken, tweetChan, rules}, nil
}

func (t *TwitterAPI) doOAuth2Request(method, url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+t.bearerToken)
	req.Header.Add("User-Agent", "CheerThemPlease")
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
	} else {
		log.Println("Adding rule success")
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
