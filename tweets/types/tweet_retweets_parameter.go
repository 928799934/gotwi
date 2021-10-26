package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type TweetRetweetsRetweetedByParams struct {
	accessToken string

	// Path parameter
	ID string // Tweet ID

	// Query parameters
	Expansions  []string
	MediaFields []string
	PlaceFields []string
	PollFields  []string
	TweetFields []string
	UserFields  []string
}

var TweetRetweetsRetweetedByQueryParams = map[string]struct{}{
	"id":           {},
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *TweetRetweetsRetweetedByParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetRetweetsRetweetedByParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetRetweetsRetweetedByParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetRetweetsRetweetedByQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetRetweetsRetweetedByParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetRetweetsRetweetedByParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.Expansions != nil && len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions)
	}

	if p.MediaFields != nil && len(p.MediaFields) > 0 {
		m["media.fields"] = util.QueryValue(p.MediaFields)
	}

	if p.PlaceFields != nil && len(p.PlaceFields) > 0 {
		m["place.fields"] = util.QueryValue(p.PlaceFields)
	}

	if p.PollFields != nil && len(p.PollFields) > 0 {
		m["poll.fields"] = util.QueryValue(p.PollFields)
	}

	if p.TweetFields != nil && len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields)
	}

	if p.UserFields != nil && len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields)
	}

	return m
}

type TweetRetweetsPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TweetID string `json:"tweet_id"`
}

func (p *TweetRetweetsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetRetweetsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetRetweetsPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *TweetRetweetsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *TweetRetweetsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type TweetRetweetsDeleteParams struct {
	accessToken string

	// Path parameter
	ID            string // The authenticated user ID
	SourceTweetID string
}

func (p *TweetRetweetsDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetRetweetsDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetRetweetsDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.SourceTweetID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.SourceTweetID)
	endpoint = strings.Replace(endpoint, ":source_tweet_id", escapedTID, 1)

	return endpoint
}

func (p *TweetRetweetsDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetRetweetsDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}