package twitter

import (
	"github.com/teitei-tk/tawawa-bot/config"

	"github.com/ChimeraCoder/anaconda"
)

type Client struct {
	APIClient *anaconda.TwitterApi
	Config    config.Twitter
}

func NewClient() (Client, error) {
	var client Client

	conf, err := config.InitTwitter()
	if err != nil {
		return client, err
	}

	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	api := anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)

	client.APIClient = api
	client.Config = conf

	return client, nil
}
