package tweettracker

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getBearerToken(consumerKey, consumerKeySecret string) (string, error) {
	payload := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest(BearerTokenURL["method"], BearerTokenURL["url"], payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return "", err
	}

	req.SetBasicAuth(consumerKey, consumerKeySecret)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Cannot get Bearer Token")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err.Error())
		return "", err
	}

	bearerTokenResponse, err := ParseBearerTokenResponse(body)
	if err != nil {
		log.Fatal("Error parsing Bearer Token JSON: ", err.Error())
		return "", err
	}
	return bearerTokenResponse.AccessToken, nil
}
