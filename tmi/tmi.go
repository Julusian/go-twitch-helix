package tmi

import "github.com/Julusian/go-twitch-helix/twitch"

const TmiURL = "https://tmi.twitch.tv/"

func newTmiRequest(path string) twitch.IRequestBuilder {
	return twitch.NewRequestBuilder(TmiURL, twitch.AuthTypeOAuth, path).
		WithAcceptHeader("application/vnd.twitchtv.v5+json")
}
