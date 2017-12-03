package helix

import (
	"encoding/json"

	"github.com/Julusian/go-twitch-helix/twitch"
)

type GamesResponse struct {
	Data       []GamesEntry `json:"data"`
	Pagination Pagination   `json:"pagination,omitempty"`
}

type GamesEntry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url,omitempty"`
}

func BuildGetGamesById(ids ...string) twitch.IRequest {
	return newHelixRequest("games").
		WithParamStringArray("id", ids).
		Get()
}
func BuildGetGamesByName(names ...string) twitch.IRequest {
	return newHelixRequest("games").
		WithParamStringArray("name", names).
		Get()
}

func GetGamesById(client *twitch.ApiClient, ids ...string) (*GamesResponse, *twitch.RateLimit, error) {
	return getGames(client, BuildGetGamesById(ids...))
}
func GetGamesByName(client *twitch.ApiClient, names ...string) (*GamesResponse, *twitch.RateLimit, error) {
	return getGames(client, BuildGetGamesByName(names...))
}
func getGames(client *twitch.ApiClient, req twitch.IRequest) (*GamesResponse, *twitch.RateLimit, error) {
	res, rate, err := client.MakeRequest(req)
	if err != nil {
		return nil, nil, err
	}

	r, err2 := UnmarshalGames(res)
	return r, rate, err2
}

func UnmarshalGames(res []byte) (*GamesResponse, error) {
	r := &GamesResponse{}
	err := json.Unmarshal(res, r)
	return r, err
}
