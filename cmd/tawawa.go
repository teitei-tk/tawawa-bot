package main

import (
	"fmt"

	"github.com/teitei-tk/tawawa-bot/twitter"
)

func main() {
	client, err := twitter.NewClient()
	if err != nil {
		panic(err)
	}

	result, err := twitter.GetAllTawawaTweets(client, twitter.RequestParametor{})
	if err != nil {
		panic(err)
	}

	for _, v := range result.Tweets {
		fmt.Println(v.Text)
		fmt.Println(twitter.FetchHTTPSMediaURL(v))
		fmt.Println("---------------------")
	}
}
