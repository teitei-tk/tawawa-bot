package line

import (
	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/teitei-tk/tawawa-bot/config"
)

type Client struct {
	APIClient *linebot.Client
}

func NewClient() (client Client, err error) {
	client = Client{}

	apiClient, err := config.InitLine()
	if err != nil {
		return client, err
	}

	return Client{
		APIClient: apiClient,
	}, nil
}
