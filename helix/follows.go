package helix

import (
	"encoding/json"
	"time"

	"github.com/Julusian/go-twitch-helix/twitch"
)

type FollowsResponse struct {
	Data       []FollowsEntry `json:"data"`
	Pagination Pagination     `json:"pagination,omitempty"`
}

type FollowsEntry struct {
	UserID     int       `json:"from_id,string"`
	ChannelID  int       `json:"to_id,string"`
	FollowedAt time.Time `json:"followed_at"`
}

type FollowsParams struct {
	After     string
	Before    string
	Limit     int
	ChannelID int
	UserID    int
}

func BuildGetFollows(params *FollowsParams) twitch.IRequest {
	return newHelixRequest("users/follows").
		WithParamString("after", params.After).
		WithParamString("before", params.Before).
		WithParamInt("first", params.Limit).
		WithParamInt("from_id", params.UserID).
		WithParamInt("to_id", params.ChannelID).
		Get()
}

func GetFollows(client *twitch.ApiClient, params *FollowsParams) (*FollowsResponse, *twitch.RateLimit, error) {
	res, rate, err := client.MakeRequest(BuildGetFollows(params))
	if err != nil {
		return nil, nil, err
	}

	r, err := UnmarshalFollows(res)
	return r, rate, err
}

func UnmarshalFollows(res []byte) (*FollowsResponse, error) {
	r := &FollowsResponse{}
	err := json.Unmarshal(res, r)
	return r, err
}
