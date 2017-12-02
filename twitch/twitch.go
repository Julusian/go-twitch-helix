package twitch

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type RateLimit struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

type ApiClient struct {
	client   *http.Client
	clientID string

	RateLimitRetries int
	DefaultAuthToken string
}

type errorBody struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	ErrorCodeUnknown        = 900
	ErrorCodeInvalidRequest = 901
	ErrorCodeApiFailure     = 910
)

type TwitchApiError struct {
	Code       int
	Message    string
	InnerError error
}

func newError(code int, message string, err error) *TwitchApiError {
	return &TwitchApiError{
		Code:       code,
		Message:    message,
		InnerError: err,
	}
}

func (e *TwitchApiError) Error() string {
	return fmt.Sprintf("%s (%d): %v", e.Message, e.Code, e.InnerError)
}

func NewApiClient(client *http.Client, clientID string) *ApiClient {
	return &ApiClient{
		client:   client,
		clientID: clientID,
	}
}

func (t *ApiClient) MakeRequest(spec IRequest) ([]byte, *RateLimit, *TwitchApiError) {
	baseURL, err := url.Parse(spec.GetBaseURL())
	if err != nil {
		return nil, nil, newError(ErrorCodeInvalidRequest, "Failed to parse request base URL", err)
	}

	// Try and parse url
	rel, err := url.Parse(spec.GetPath())
	if err != nil {
		return nil, nil, newError(ErrorCodeInvalidRequest, "Failed to parse request path", err)
	}

	// Ensure the request doesnt hit a cache
	spec.GetQueryParams()["_"] = []string{strconv.Itoa(int(time.Now().Unix()))}

	rel.RawQuery = url.Values(spec.GetQueryParams()).Encode()
	u := baseURL.ResolveReference(rel)

	// Create request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, newError(ErrorCodeInvalidRequest, "Failed to build http request from url", err)
	}

	if spec.GetAcceptHeader() != "" {
		req.Header.Add("Accept", spec.GetAcceptHeader())
	}

	if !spec.GetRemoveClientID() {
		req.Header.Add("Client-ID", t.clientID)

		// Add oauth token if supplied
		token := spec.GetAuthToken()
		if token == "" {
			token = t.DefaultAuthToken
		}
		if token != "" {
			switch spec.GetAuthType() {
			case AuthTypeOAuth:
				req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", token))
			case AuthTypeBearer:
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			}
		}
	}

	return t.runRequest(req, t.RateLimitRetries)
}

func (t *ApiClient) runRequest(req *http.Request, retries int) ([]byte, *RateLimit, *TwitchApiError) {
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, nil, newError(ErrorCodeUnknown, "Unknown http error", err)
	}

	defer resp.Body.Close()
	rateLimit := parseRateLimitHeaders(resp.Header)

	switch resp.StatusCode {
	case http.StatusOK:
		fallthrough
	case http.StatusNotModified:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, rateLimit, newError(ErrorCodeApiFailure, "Failed to read response body", err)
		}

		return body, rateLimit, nil
	case http.StatusTooManyRequests:
		if retries <= 0 {
			return nil, rateLimit, newError(http.StatusTooManyRequests, "Rate limited. No more retries", nil)
		}

		// Wait until rate limit resets
		time.Sleep(15 * time.Second)                                  // Wait for some period. Note: the counter will change before the reset time
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond) // add some jitter

		return t.runRequest(req, retries-1)
	case http.StatusServiceUnavailable:

		// TODO - wait and retry before failing?
		return nil, rateLimit, newError(ErrorCodeApiFailure, "Service unavailable", nil)
	case http.StatusBadRequest:
		fallthrough
	default:
		return nil, rateLimit, tryParseResponse(resp.StatusCode, resp.Body)
	}
}

func tryParseResponse(statusCode int, body io.ReadCloser) *TwitchApiError {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return newError(ErrorCodeApiFailure, "Failed to read response", err)
	}

	res := &errorBody{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return newError(ErrorCodeApiFailure, "Failed to decode response", err)
	}

	return newError(res.Status, fmt.Sprintf("%s: %s", res.Error, res.Message), nil)
}

func parseRateLimitHeaders(headers http.Header) *RateLimit {
	rateLimit := &RateLimit{}

	limit, err := strconv.Atoi(headers.Get("Ratelimit-Limit"))
	if err == nil {
		rateLimit.Limit = limit
	}

	remain, err := strconv.Atoi(headers.Get("Ratelimit-Remaining"))
	if err == nil {
		rateLimit.Remaining = remain
	}

	reset, err := strconv.Atoi(headers.Get("Ratelimit-Reset"))
	if err == nil {
		rateLimit.Reset = time.Unix(int64(reset), 0)
	}

	return rateLimit
}
