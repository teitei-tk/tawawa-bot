package twitter

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/teitei-tk/tawawa-bot/config"
)

const (
	// Tawawa Owner Name. from https://twitter.com/Strangestone/
	OwnerScreenName = "Strangestone"

	// from https://twitter.com/Strangestone/status/569617644472573952
	sinceTweetID             = 569617644472573952
	defaultFindTimelineCount = 200

	TwitterResponseCacheKey = "Tawawa"
	ResponseCacheExpiredAt  = 10000
)

type RequestParametor struct {
	ScreenName      string
	Count           int
	IncludeReTweets bool
	SinceID         int64
	MaxID           int64
}

// Get Tawawa Owner Timeline it's use Cache
func GetOwnerTimeline(client Client, param RequestParametor) (res UserTimelineResponse, err error) {
	res = UserTimelineResponse{}

	redisClient := config.InitRedis()
	val, err := redisClient.Get(TwitterResponseCacheKey).Result()
	if val != "" {
		err = json.Unmarshal([]byte(val), &res)
		if err != nil {
			return res, err
		}
		return res, err
	}

	res, err = getOwnerTimelineFromTwitter(client, param)
	if err != nil {
		return res, err
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		return res, err
	}

	err = redisClient.Set(TwitterResponseCacheKey, string(bytes), 0).Err()
	if err != nil {
		return res, err
	}

	return res, err
}

func getOwnerTimelineFromTwitter(client Client, param RequestParametor) (res UserTimelineResponse, err error) {
	res = UserTimelineResponse{}

	// merge
	mergedParam := RequestParametor{
		ScreenName:      OwnerScreenName,
		Count:           defaultFindTimelineCount,
		IncludeReTweets: false,
		SinceID:         sinceTweetID,
		MaxID:           param.MaxID,
	}

	values := make(url.Values)
	values.Add("screen_name", mergedParam.ScreenName)
	values.Add("count", strconv.Itoa(mergedParam.Count))
	values.Add("include_rts", strconv.FormatBool(mergedParam.IncludeReTweets))
	values.Add("since_id", strconv.FormatInt(mergedParam.SinceID, 10))

	if param.MaxID != 0 {
		values.Add("max_id", strconv.FormatInt(mergedParam.MaxID, 10))
	}

	results, err := client.APIClient.GetUserTimeline(values)
	if err != nil {
		return res, err
	}

	res.Tweets = results
	return res, nil
}
