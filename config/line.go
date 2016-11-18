package config

import (
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func InitLine() (client *linebot.Client, err error) {
	lineClient, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		return lineClient, err
	}

	return lineClient, nil
}
