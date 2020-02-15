package tweettracker

import (
	"log"

	"github.com/dghubble/sling"
)

func getBearerToken(consumerKey, consumerKeySecret string) (string, error) {
	payload := &BearerTokenRequest{GrantType: "client_credentials"}

	bearerTokenResponse := new(BearerTokenResponse)
	resp, err := sling.New().
		Post(BearerTokenURL["url"]).
		BodyForm(payload).
		SetBasicAuth(consumerKey, consumerKeySecret).
		Set("User-Agent", "CheerMeUpPlease").
		ReceiveSuccess(&bearerTokenResponse)

	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Cannot get Bearer Token")
	} else {
		log.Println("Get bearer token success")
	}

	return bearerTokenResponse.AccessToken, nil
}
