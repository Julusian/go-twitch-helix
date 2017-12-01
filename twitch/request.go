package twitch

var _ IRequest = (*request)(nil)

type IRequest interface {
	GetBaseURL() string
	GetPath() string
	GetQueryParams() map[string][]string

	GetOAuthToken() string
	GetBearerToken() string
	GetAcceptHeader() string
}

type request struct {
	BaseURL     string
	Path        string
	QueryParams map[string][]string

	OAuthToken   string
	BearerToken  string
	AcceptHeader string
}

func (r *request) GetBaseURL() string {
	return r.BaseURL
}

func (r *request) GetOAuthToken() string {
	return r.OAuthToken
}

func (r *request) GetBearerToken() string {
	return r.BearerToken
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
