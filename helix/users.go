package helix

import (
	"encoding/json"

	"github.com/julusian/go-twitch-helix/twitch"
)

type UserType string
type BroadcasterType string

const (
	UserTypeStaff     UserType = "staff"
	UserTypeAdmin     UserType = "admin"
	UserTypeGlobalMod UserType = "global_mod"
	UserTypeNormal    UserType = ""

	BroadcasterTypePartner   UserType = "partner"
	BroadcasterTypeAffiliate UserType = "affiliate"
	BroadcasterTypeNormal    UserType = ""
)

type UsersResponse struct {
	Data       []UsersEntry `json:"data"`
	Pagination Pagination   `json:"pagination,omitempty"`
}

type UsersEntry struct {
	BroadcasterType BroadcasterType `json:"broadcaster_type"`
	Description     string          `json:"description"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ID              int             `json:"id,string"`
	Login           string          `json:"login"`
	OfflineImageURL string          `json:"offline_image_url"`
	ProfileImageURL string          `json:"profile_image_url"`
	Type            UserType        `json:"type"`
	ViewCount       int             `json:"view_count"`
}

type UsersParams struct {
	IDs       []int
	Logins    []string
	AuthToken string
}

func BuildGetUsers(params *UsersParams) twitch.IRequest {
	return newHelixRequest("users").
		WithBearerToken(params.AuthToken).
		WithParamIntArray("id", params.IDs).
		WithParamStringArray("login", params.Logins).
		Get()
}

func GetUsers(client *twitch.ApiClient, params *UsersParams) (*UsersResponse, *twitch.RateLimit, error) {
	res, rate, err := client.MakeRequest(BuildGetUsers(params))
	if err != nil {
		return nil, nil, err
	}

	r, err := UnmarshalUsers(res)
	return r, rate, err
}

func UnmarshalUsers(res []byte) (*UsersResponse, error) {
	r := &UsersResponse{}
	err := json.Unmarshal(res, r)
	return r, err
}
