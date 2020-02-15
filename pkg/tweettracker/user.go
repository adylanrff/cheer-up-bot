package tweettracker

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

func (t *TwitterAPI) GetUser(userId string) (*User, error) {
	payload := GetUserRequest{UserID: []string{userId}}
	resp, err := t.doOAuth1Request(GetUserInfoURL["method"], GetUserInfoURL["url"], payload)

	if err != nil {
		log.Fatal("Error getting user info: ", err)
		return nil, err
	}

	var users []User

	if resp.StatusCode != 200 {
		log.Fatal("Failed fetching user: ", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &users)

	if err != nil {
		log.Fatal("Error marshalling user info: ", err)
		return nil, err
	}

	return &users[0], nil
}
