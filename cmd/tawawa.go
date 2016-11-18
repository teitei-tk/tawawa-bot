package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/teitei-tk/tawawa-bot/line"
	"github.com/teitei-tk/tawawa-bot/twitter"
	"github.com/teitei-tk/tawawa-bot/util"
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
			replySource := util.GetReplySource(event)
			if event.Type == linebot.EventTypeJoin || event.Type == linebot.EventTypeFollow {
				textMsg := linebot.NewTextMessage("たわわをおくれ と言ってみましょう。")
				lineClient.APIClient.PushMessage(replySource, textMsg).Do()
				return
			}

			if event.Type != linebot.EventTypeMessage {
				continue
			}

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "たわわをおくれ" {
					tweet := line.RandResponceChoice(res)
					mediaURL := twitter.FetchHTTPSMediaURL(tweet)

					textMsg := linebot.NewTextMessage(util.ToLowerString(tweet.Text))
					imageMsg := linebot.NewImageMessage(mediaURL, mediaURL)

					if _, err := lineClient.APIClient.PushMessage(replySource, textMsg, imageMsg).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}

		return
	})

	err = http.ListenAndServeTLS(":"+os.Getenv("LISTEN_PORT"), os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"), nil)
	if err != nil {
		panic(err)
	}

}
