package helix

import (
	"encoding/json"
	"time"

	"github.com/julusian/go-twitch-helix/twitch"
)

type StreamsResponse struct {
	Data       []StreamsEntry `json:"data"`
	Pagination Pagination     `json:"pagination,omitempty"`
}

type StreamsEntry struct {
	CommunityIDs []string  `json:"community_ids"`
	GameID       string    `json:"game_id"`
	ID           int       `json:"id,string"`
	Language     string    `json:"language"`
	StartedAt    time.Time `json:"started_at"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	UserID       int       `json:"user_id,string"`
	ViewerCount  int       `json:"viewer_count"`
}

type StreamsParams struct {
	After       string
	Before      string
	CommunityID []string
	Limit       int
	GameID      []string
	Language    []string
	Type        string // all, live, vodcast
	UserID      []int
	UserLogin   []string
}

func BuildGetStreams(params *StreamsParams) twitch.IRequest {
	return newHelixRequest("streams").
		WithParamString("after", params.After).
		WithParamString("before", params.Before).
		WithParamStringArray("community_id", params.CommunityID).
		WithParamInt("first", params.Limit).
		WithParamStringArray("game_id", params.GameID).
		WithParamStringArray("language", params.Language).
		WithParamString("type", params.Type).
		WithParamIntArray("user_id", params.UserID).
		WithParamStringArray("user_login", params.UserLogin).
		Get()
}

func GetStreams(client *twitch.ApiClient, params *StreamsParams) (*StreamsResponse, *twitch.RateLimit, error) {
	res, rate, err := client.MakeRequest(BuildGetStreams(params))
	if err != nil {
		return nil, nil, err
	}

	r, err := UnmarshalStreams(res)
	return r, rate, err
}

func UnmarshalStreams(res []byte) (*StreamsResponse, error) {
	r := &StreamsResponse{}
	err := json.Unmarshal(res, r)
	return r, err
}
