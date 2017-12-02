package twitch

import "encoding/json"

var _ IRequest = (*request)(nil)

type IRequest interface {
	GetBaseURL() string
	GetPath() string
	GetQueryParams() map[string][]string

	GetRemoveClientID() bool
	GetAuthToken() string
	GetAuthType() AuthType
	GetAcceptHeader() string
}

type AuthType int

const (
	AuthTypeOAuth  AuthType = 0
	AuthTypeBearer AuthType = 1
)

type request struct {
	BaseURL     string
	Path        string
	QueryParams map[string][]string

	RemoveClientID bool
	AuthType       AuthType
	AuthToken      string `json:",omitempty"`
	AcceptHeader   string `json:",omitempty"`
}

func (r *request) GetBaseURL() string {
	return r.BaseURL
}

func (r *request) GetRemoveClientID() bool {
	return r.RemoveClientID
}

func (r *request) GetAuthToken() string {
	return r.AuthToken
}

func (r *request) GetAuthType() AuthType {
	return r.AuthType
}

func (r *request) GetAcceptHeader() string {
	return r.AcceptHeader
}

func (r *request) GetPath() string {
	return r.Path
}

func (r *request) GetQueryParams() map[string][]string {
	return r.QueryParams
}

func UnmarshalRequest(data []byte) (IRequest, error) {
	res := request{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
