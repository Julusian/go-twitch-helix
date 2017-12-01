package helix

import (
	"net/http"
	"os"
	"testing"

	"github.com/julusian/go-twitch-helix/twitch"
)

func newTestClient(t *testing.T) *twitch.ApiClient {
	return twitch.NewApiClient(&http.Client{}, os.Getenv("TWITCH_CLIENTID"))
}
