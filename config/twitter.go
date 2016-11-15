package config

import (
	"errors"
	"os"
)

type Twitter struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func InitTwitter() (Twitter, error) {
	var tw Twitter

	key := os.Getenv("TWITTER_CONSUMER_KEY")
	secrect := os.Getenv("TWITTER_CONSUMER_SECRET")
	if key == "" || secrect == "" {
		return tw, errors.New("ConsumerKey or ConsumerSecret is not Defined.")
	}

	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accesssTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	if accessToken == "" || accesssTokenSecret == "" {
		return tw, errors.New("AccessToken or AccesssTokenSecret is Not Defined.")
	}

	tw = Twitter{
		ConsumerKey:       key,
		ConsumerSecret:    secrect,
		AccessToken:       accessToken,
		AccessTokenSecret: accesssTokenSecret,
	}

	return tw, nil
}
