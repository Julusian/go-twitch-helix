package helix

import (
	"net/http"
	"os"
	"testing"

	"github.com/Julusian/go-twitch-helix/twitch"
)

func newTestClient(t *testing.T) *twitch.ApiClient {
	cl := twitch.NewApiClient(&http.Client{}, os.Getenv("TWITCH_CLIENTID"))
	cl.RateLimitRetries = 2
	return cl
}
