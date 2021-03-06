package twitter

import (
	"io"
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
	tweetChan   chan *Tweet
	rules       []TweetFilterRule
}

func NewTwitterAPI(config *TwitterConfig, handler TwitterHandler, rules []TweetFilterRule) (*TwitterAPI, error) {
	bearerToken, err := getBearerToken(config.APIKey, config.APISecretKey)
	if err != nil {
		log.Fatal("Error getting bearer token: ", err.Error())
		return nil, err
	}

	tweetChan := make(chan *Tweet, 10)
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

	var req *http.Request
	var err error
	s := sling.New().QueryStruct(payload)

	switch method {
	case "GET":
		req, err = s.Get(url).Request()
	case "POST":
		req, err = s.Post(url).Request()
	}

	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return nil, err
	}

	resp, err := httpClient.Do(req)
	return resp, nil
}

func (t *TwitterAPI) ResetRules() {

	// Get current rules
	var ruleValues []string
	rules, err := t.GetRules()
	if err != nil {
		log.Panicln("Error getting rules: ", err)
	}
	for _, rule := range rules {
		ruleValues = append(ruleValues, rule.Value)
	}

	// Reset rules
	err = t.DeleteAllRules(ruleValues)
	if err != nil {
		log.Fatal("Error resetting rules: ", err)
	}
	err = t.AddRule()
	if err != nil {
		log.Fatal("Error adding rule: ", err)
	}

	newRules, err := t.GetRules()
	if err != nil {
		log.Panicln("Error getting rules: ", err)
	}
	log.Println("Filter rules: ", newRules)

}

func (t *TwitterAPI) Run() {
	t.ResetRules()
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

	// goroutine for handling tweets
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
