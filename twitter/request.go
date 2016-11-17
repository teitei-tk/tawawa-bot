package twitter

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"

	"github.com/teitei-tk/tawawa-bot/config"
)

const (
	// Tawawa Owner Name. from https://twitter.com/Strangestone/
	OwnerScreenName = "Strangestone"

	// from https://twitter.com/Strangestone/status/569617644472573952
	sinceTweetID             = 569617644472573952
	defaultFindTimelineCount = 200

	TwitterResponseCacheKey = "Tawawa"
	ResponseCacheExpiredAt  = 1000000000
)

type RequestParametor struct {
	ScreenName      string
	Count           int
	IncludeReTweets bool
	SinceID         int64
	MaxID           int64
}

// Get Tawawa Owner Timeline it's use Cache
func GetAllTawawaTweets(client Client, param RequestParametor) (res UserTimelineResponse, err error) {
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

	result, err := FetchAllTawawaTweets(client, param)
	if err != nil {
		return res, err
	}
	res.Tweets = result

	bytes, err := json.Marshal(res)
	if err != nil {
		return res, err
	}

	err = redisClient.Set(TwitterResponseCacheKey, string(bytes), time.Duration(time.Millisecond*ResponseCacheExpiredAt)).Err()
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
		SinceID:         param.SinceID,
		MaxID:           param.MaxID,
	}

	values := make(url.Values)
	values.Add("screen_name", mergedParam.ScreenName)
	values.Add("count", strconv.Itoa(mergedParam.Count))
	values.Add("include_rts", strconv.FormatBool(mergedParam.IncludeReTweets))
	if param.SinceID != 0 {
		values.Add("since_id", strconv.FormatInt(mergedParam.SinceID, 10))
	} else {
		values.Add("since_id", strconv.FormatInt(sinceTweetID, 10))
	}

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

func showFromOwnerTweet(tweetID int64, client Client) (resTweet anaconda.Tweet, err error) {
	var tweet anaconda.Tweet

	values := make(url.Values)
	values.Add("id", strconv.FormatInt(tweetID, 10))
	values.Add("include_entities", strconv.FormatBool(true))

	res, err := client.APIClient.GetTweet(tweetID, values)
	if err != nil {
		return tweet, err
	}

	return res, nil
}

func FetchAllTawawaTweets(client Client, param RequestParametor) (filterdTweets []anaconda.Tweet, err error) {
	var maxRequestTweetID int64
	var sinceID int64
	var tweestCount int
	var tweets []anaconda.Tweet

	for {
		if maxRequestTweetID != 0 {
			param.MaxID = maxRequestTweetID
		}

		res, err := getOwnerTimelineFromTwitter(client, param)
		tweestCount = tweestCount + len(res.Tweets)
		if err != nil {
			return tweets, err
		}

		filterdRes := FilterTawawaTweets(res)
		for _, t := range filterdRes.Tweets {
			if maxRequestTweetID == 0 || maxRequestTweetID > t.Id {
				maxRequestTweetID = t.Id
			}

			if sinceID == 0 || t.Id > sinceID {
				sinceID = t.Id
			}

			tweets = append(tweets, t)
		}

		if maxRequestTweetID == sinceTweetID || tweestCount >= 3200 {
			break
		}
	}

	var showTweetRequestCount = 0
	var requestTweetID = maxRequestTweetID
	for {
		resTweet, err := showFromOwnerTweet(requestTweetID, client)
		if err != nil {
			return tweets, err
		}

		tweets = append(tweets, resTweet)

		requestTweetID = resTweet.InReplyToStatusID
		showTweetRequestCount = showTweetRequestCount + 1
		if showTweetRequestCount >= 100 || resTweet.Id == sinceTweetID {
			break
		}
	}

	return tweets, nil
}
