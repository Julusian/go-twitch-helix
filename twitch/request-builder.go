package twitch

import "strconv"

var _ IRequestBuilder = (*requestBuilder)(nil)

func NewRequestBuilder(baseURL string, authType AuthType, path string) IRequestBuilder {
	return &requestBuilder{
		req: &request{
			BaseURL:     baseURL,
			Path:        path,
			QueryParams: map[string][]string{},
			AuthType:    authType,
		},
	}
}

type IRequestBuilder interface {
	Get() IRequest

	WithAuthToken(string) IRequestBuilder
	WithAcceptHeader(string) IRequestBuilder
	WithoutClientID() IRequestBuilder

	WithParamString(name string, value string) IRequestBuilder
	WithParamStringArray(name string, value []string) IRequestBuilder
	WithParamInt(name string, value int) IRequestBuilder
	WithParamIntArray(name string, value []int) IRequestBuilder
}

type requestBuilder struct {
	req *request
}

func (b *requestBuilder) Get() IRequest {
	return b.req
}

func (b *requestBuilder) WithAuthToken(token string) IRequestBuilder {
	b.req.AuthToken = token
	return b
}

func (b *requestBuilder) WithAcceptHeader(token string) IRequestBuilder {
	b.req.AcceptHeader = token
	return b
}

func (b *requestBuilder) WithoutClientID() IRequestBuilder {
	b.req.RemoveClientID = true
	return b
}

func (b *requestBuilder) WithParamString(name string, value string) IRequestBuilder {
	if value != "" {
		b.req.QueryParams[name] = []string{value}
	}
	return b
}

func (b *requestBuilder) WithParamStringArray(name string, value []string) IRequestBuilder {
	if len(value) != 0 {
		b.req.QueryParams[name] = value
	}
	return b
}

func (b *requestBuilder) WithParamInt(name string, value int) IRequestBuilder {
	if value != 0 {
		b.req.QueryParams[name] = []string{strconv.Itoa(value)}
	}
	return b
}

func (b *requestBuilder) WithParamIntArray(name string, value []int) IRequestBuilder {
	if len(value) != 0 {
		strVals := make([]string, len(value))
		for i, v := range value {
			strVals[i] = strconv.Itoa(v)
		}
		b.req.QueryParams[name] = strVals
	}
	return b
}
