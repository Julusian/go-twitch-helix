package helix

import "github.com/Julusian/go-twitch-helix/twitch"

const HelixURL = "https://api.twitch.tv/helix/"

func newHelixRequest(path string) twitch.IRequestBuilder {
	return twitch.NewRequestBuilder(HelixURL, twitch.AuthTypeBearer, path)
}
