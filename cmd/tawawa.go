package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/teitei-tk/tawawa-bot/line"
	"github.com/teitei-tk/tawawa-bot/twitter"
)

func main() {
	client, err := twitter.NewClient()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		res, err := twitter.GetAllTawawaTweets(client, twitter.RequestParametor{})
		if err != nil {
			panic(err)
		}

		lineClient, err := line.NewClient()
		events, err := lineClient.APIClient.ParseRequest(req)
		if err != nil {
			panic(err)
		}

		for _, event := range events {
			if event.Type != linebot.EventTypeMessage {
				continue
			}

			tweet := line.RandResponceChoice(res)
			content := line.FetchLineImages(tweet)

			msg := linebot.NewImageMessage(content.ContentURL, content.DisplayURL)
			if _, err = lineClient.APIClient.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
				log.Print(err)
			}
		}

		return
	})

	err = http.ListenAndServeTLS(":"+os.Getenv("LISTEN_PORT"), os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"), nil)
	if err != nil {
		panic(err)
	}

}
