package twitter

import (
	"strings"

	"github.com/ChimeraCoder/anaconda"

	"github.com/teitei-tk/tawawa-bot/util"
)

var (
	TawawaTweetText = []string{"月曜日のたわわ", "月曜朝の社畜諸兄にたわわをお届けします"}
)

// Twitter UserTimeline Response. doc at https://dev.twitter.com/rest/reference/get/statuses/user_timeline
type UserTimelineResponse struct {
	Tweets []anaconda.Tweet `json:"tweets"`
}

type FilterTawawaResponse struct {
	Tweets []anaconda.Tweet
}

func FetchHTTPSMediaURL(tweet anaconda.Tweet) (url string) {
	var mediaURL string

	for _, media := range tweet.Entities.Media {
		if media.Media_url_https == "" {
			continue
		}

		mediaURL = media.Media_url_https
		break
	}

	return mediaURL
}

func HasURLInTweet(tweet anaconda.Tweet) bool {
	return strings.TrimSpace(FetchHTTPSMediaURL(tweet)) != ""
}

func IsTawawaTweet(tweet anaconda.Tweet) bool {
	isTawawaTweet := false
	for _, tawawaText := range TawawaTweetText {
		if HasURLInTweet(tweet) == false {
			continue
		}

		if strings.Index(tweet.Text, tawawaText) == -1 {
			continue
		}

		if util.IsTawawaString(tweet.Text) {
			isTawawaTweet = true
		}
	}

	return isTawawaTweet
}

func FilterTawawaTweets(timelime UserTimelineResponse) (res FilterTawawaResponse) {
	filterdTawawaResponse := FilterTawawaResponse{}

	var filterdTweets []anaconda.Tweet
	for _, tweet := range timelime.Tweets {
		if IsTawawaTweet(tweet) {
			filterdTweets = append(filterdTweets, tweet)
		}
	}

	filterdTawawaResponse.Tweets = filterdTweets
	return filterdTawawaResponse
}
