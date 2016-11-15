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

	res, err := twitter.GetOwnerTimeline(client, twitter.RequestParametor{})
	if err != nil {
		panic(err)
	}

	for _, v := range res.Tweets {
		fmt.Println(v.Text)
	}
}