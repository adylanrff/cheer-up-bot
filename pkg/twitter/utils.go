package twitter

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
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
	}

	return bearerTokenResponse.AccessToken, nil
}

func convertImageToBase64String(img image.Image) (string, error) {
	var base64img []byte
	buf := new(bytes.Buffer)
	imgBytes := buf.Bytes()
	err := png.Encode(buf, img)

	if err != nil {
		log.Fatal("PNG encoding failed", err)
		return "", err
	}

	base64.StdEncoding.Encode(imgBytes, base64img)
	return string(base64img), nil
}
