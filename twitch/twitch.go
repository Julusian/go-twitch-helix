package twitch

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
}

type errorBody struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewApiClient(client *http.Client, clientID string) *ApiClient {
	return &ApiClient{
		client:   client,
		clientID: clientID,
	}
}

func (t *ApiClient) MakeRequest(spec IRequest) ([]byte, *RateLimit, error) {
	baseURL, err := url.Parse(spec.GetBaseURL())
	if err != nil {
		return nil, nil, err // TODO - wrap better
	}

	// Try and parse url
	rel, err := url.Parse(spec.GetPath())
	if err != nil {
		return nil, nil, err // TODO - wrap better
	}

	rel.RawQuery = url.Values(spec.GetQueryParams()).Encode()
	u := baseURL.ResolveReference(rel)

	// Create request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	if spec.GetAcceptHeader() != "" {
		req.Header.Add("Accept", spec.GetAcceptHeader())
	}
	req.Header.Add("Client-ID", t.clientID)

	// Add oauth token if supplied
	if spec.GetOAuthToken() != "" {
		req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", spec.GetOAuthToken()))
	}
	if spec.GetBearerToken() != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", spec.GetBearerToken()))
	}

	return t.runRequest(req, t.RateLimitRetries)
}

func (t *ApiClient) runRequest(req *http.Request, retries int) ([]byte, *RateLimit, error) {
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	rateLimit := parseRateLimitHeaders(resp.Header)

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		fallthrough
	case http.StatusNotModified:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, rateLimit, err
		}

		return body, rateLimit, nil
	case http.StatusTooManyRequests:
		if retries <= 0 {
			return nil, rateLimit, fmt.Errorf("Exceeded rate limiter retries") // TODO - better
		}

		time.Sleep(rateLimit.Reset.Sub(time.Now()))

		return t.runRequest(req, retries-1)
	case http.StatusServiceUnavailable:
		// TODO - wait and retry before failing?
		return nil, rateLimit, fmt.Errorf("Service unavailable") // TODO - better
	case http.StatusBadRequest:
		fallthrough
	default:
		return nil, rateLimit, tryParseResponse(resp.StatusCode, resp.Body)
	}
}

func tryParseResponse(statusCode int, body io.ReadCloser) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("Api error, response code: %d", statusCode)
	}

	res := &errorBody{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return fmt.Errorf("Api error, response code: %d", statusCode)
	}

	return fmt.Errorf("Twitch error: %s: %s", res.Error, res.Message)
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
