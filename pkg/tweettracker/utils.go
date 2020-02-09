package tweettracker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getBearerToken(consumerKey, consumerKeySecret string) (string, error) {
	payload := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest("POST", BearerTokenURL, payload)

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err.Error())
		return "", err
	}

	fmt.Println(string(body))

	bearerTokenResponse, err := ParseBearerTokenResponse(body)
	if err != nil {
		log.Fatal("Error parsing Bearer Token JSON: ", err.Error())
		return "", err
	}
	return bearerTokenResponse.AccessToken, nil
}
