package helix

import (
	"encoding/json"
	"time"

	"github.com/Julusian/go-twitch-helix/twitch"
)

type VideoPeriod string
type VideoSort string
type VideoType string

const (
	VideoPeriodAll   VideoPeriod = "all"
	VideoPeriodDay   VideoPeriod = "day"
	VideoPeriodMonth VideoPeriod = "month"
	VideoPeriodWeek  VideoPeriod = "week"

	VideoSortTime     VideoSort = "time"
	VideoSortTrending VideoSort = "trending"
	VideoSortViews    VideoSort = "views"

	VideoTypeAll       VideoType = "all"
	VideoTypeUpload    VideoType = "upload"
	VideoTypeArchive   VideoType = "archive"
	VideoTypeHighlight VideoType = "highlight"
)

type VideosResponse struct {
	Data       []VideosEntry `json:"data"`
	Pagination Pagination    `json:"pagination,omitempty"`
}

type VideosEntry struct {
	ID           int       `json:"id,string"`
	UserID       int       `json:"user_id,string"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	PublishedAt  time.Time `json:"published_at"`
	ThumbnailURL string    `json:"thumbnail_url"`
	ViewCount    int       `json:"view_count"`
	Language     string    `json:"language"`
	Duration     string    `json:"duration"`
}

type VideosParams struct {
	After    string
	Before   string
	Limit    int
	Language string
	Period   VideoPeriod
	Sort     VideoSort
	Type     VideoType
}

func BuildGetVideosById(ids ...int) twitch.IRequest {
	return newHelixRequest("videos").
		WithParamIntArray("id", ids).
		Get()
}

func BuildGetVideosByUserId(userID int, params *VideosParams) twitch.IRequest {
	req := newHelixRequest("videos").
		WithParamInt("user_id", userID)

	return addVideoParams(req, params).Get()
}

func BuildGetVideosByGameId(gameID string, params *VideosParams) twitch.IRequest {
	req := newHelixRequest("videos").
		WithParamString("game_id", gameID)

	return addVideoParams(req, params).Get()
}

func addVideoParams(req twitch.IRequestBuilder, params *VideosParams) twitch.IRequestBuilder {
	return req.WithParamString("after", params.After).
		WithParamString("before", params.Before).
		WithParamInt("first", params.Limit).
		WithParamString("language", params.Language).
		WithParamString("period", string(params.Period)).
		WithParamString("sort", string(params.Sort)).
		WithParamString("type", string(params.Type))
}

func GetVideosById(client *twitch.ApiClient, ids ...int) (*VideosResponse, *twitch.RateLimit, error) {
	return getVideos(client, BuildGetVideosById(ids...))
}
func GetVideosByUserId(client *twitch.ApiClient, userID int, params *VideosParams) (*VideosResponse, *twitch.RateLimit, error) {
	return getVideos(client, BuildGetVideosByUserId(userID, params))
}
func GetVideosByGameId(client *twitch.ApiClient, gameID string, params *VideosParams) (*VideosResponse, *twitch.RateLimit, error) {
	return getVideos(client, BuildGetVideosByGameId(gameID, params))
}

func getVideos(client *twitch.ApiClient, req twitch.IRequest) (*VideosResponse, *twitch.RateLimit, error) {
	res, rate, err := client.MakeRequest(req)
	if err != nil {
		return nil, nil, err
	}

	r, err2 := UnmarshalVideos(res)
	return r, rate, err2
}

func UnmarshalVideos(res []byte) (*VideosResponse, error) {
	r := &VideosResponse{}
	err := json.Unmarshal(res, r)
	return r, err
}
