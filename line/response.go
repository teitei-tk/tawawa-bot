package line

import (
	"math/rand"

	"github.com/ChimeraCoder/anaconda"

	"github.com/teitei-tk/tawawa-bot/twitter"
)

func RandResponceChoice(timeline twitter.UserTimelineResponse) (tweet anaconda.Tweet) {
	return timeline.Tweets[rand.Intn(len(timeline.Tweets))]
}
