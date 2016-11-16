package twitter

import (
	"github.com/ChimeraCoder/anaconda"
)

// Twitter UserTimeline Response. doc at https://dev.twitter.com/rest/reference/get/statuses/user_timeline
type UserTimelineResponse struct {
	Tweets []anaconda.Tweet `json:"tweets"`
}

type FilterTawawaResponse struct {
	Tweets []anaconda.Tweet
}
