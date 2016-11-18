package line

import (
	"math/rand"

	"github.com/ChimeraCoder/anaconda"

	"github.com/teitei-tk/tawawa-bot/twitter"
)

type TawawaImageResponse struct {
	ContentURL string
	DisplayURL string
}

func RandResponceChoice(timeline twitter.UserTimelineResponse) (tweet anaconda.Tweet) {
	return timeline.Tweets[rand.Intn(len(timeline.Tweets))]
}

func FetchLineImages(tweet anaconda.Tweet) (conent TawawaImageResponse) {
	var content = TawawaImageResponse{}

	for _, media := range tweet.Entities.Media {
		if media.Media_url_https != "" {
			content.ContentURL = media.Media_url_https
			content.DisplayURL = media.Display_url
		}
	}

	return content
}
