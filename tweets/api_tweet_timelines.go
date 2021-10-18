package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetTimelinesTweetsEndpoint = "https://api.twitter.com/2/users/:id/tweets"
)

// Returns Tweets composed by a single user, specified by the requested user ID.
// By default, the most recent ten Tweets are returned per request. Using pagination, the most recent 3,200 Tweets can be retrieved.
// The Tweets returned by this endpoint count towards the Project-level Tweet cap.
// https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
func TweetTimelinesTweets(c *gotwi.GotwiClient, p *types.TweetTimelinesTweetsParams) (*types.TweetTimelinesTweetsResponse, error) {
	res := &types.TweetTimelinesTweetsResponse{}
	if err := c.CallAPI(TweetTimelinesTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
