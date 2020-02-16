package twitter

import (
	"encoding/json"
	"image"
	"io/ioutil"
	"log"
)

type Media struct {
	MediaID          string `json:"media_id_string"`
	MediaKey         string `json:"media_key"`
	Size             int    `json:"size"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	Image            Image  `json:"image"`
}

type Image struct {
	ImageType string `json:"image_type"`
	Width     int    `json:"w"`
	Height    int    `json:"h"`
}

func (t *TwitterAPI) UploadImage(img image.Image) (*Media, error) {
	base64img, err := convertImageToBase64String(img)
	if err != nil {
		log.Fatal("Error converting image to Base64")
		return nil, err
	}

	request := &UploadMediaRequest{
		MediaData:     base64img,
		MediaCategory: "tweet_image",
	}

	resp, err := t.doOAuth1Request(UploadImageURL["method"], UploadImageURL["url"], request)

	if err != nil || resp.StatusCode != 200 {
		log.Fatal("Error uploading image")
		return nil, err
	}

	var media Media
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &media)

	if err != nil {
		log.Fatal("Error unmarshalling media json response")
		return nil, err
	}

	return &media, nil
}
