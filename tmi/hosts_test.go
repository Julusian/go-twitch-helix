package tmi

import (
	"testing"

	"github.com/Julusian/go-twitch-helix/helix"
	"github.com/Julusian/go-twitch-helix/twitch"
	"github.com/stretchr/testify/require"
)

func TestHosts(t *testing.T) {
	tc := newTestClient(t)

	params := &helix.StreamsParams{
		Limit: 10,
	}

	res, _, err := helix.GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	for _, v := range res.Data {
		if testHostsForChannel(t, tc, v.UserID) {
			return
		}
	}

	require.Fail(t, "Expected one of the channels to pass")
}

func testHostsForChannel(t *testing.T, tc *twitch.ApiClient, channel int) bool {
	res, _, err := GetHosts(tc, channel)
	if err != nil {
		err2 := err.(*twitch.TwitchApiError)
		require.Equal(t, err2.Code, twitch.ErrorCodeApiFailure)
		require.Equal(t, err2.Message, "Failed to decode response")

		// Try one more time
		res, _, err = GetHosts(tc, channel)
		require.Nil(t, err)
		require.NotNil(t, res)
	}
	require.NotNil(t, res)

	if len(res.Hosts) < 5 { // Need a channel with a good sample
		return false
	}

	oneWasPartnered := false
	for _, v := range res.Hosts {
		require.NotEqual(t, 0, v.HostID)
		require.NotEqual(t, "", v.HostLogin)
		require.NotEqual(t, "", v.HostDisplayName)
		require.NotEqual(t, 0, v.TargetID)
		require.NotEqual(t, "", v.TargetLogin)
		require.NotEqual(t, "", v.TargetDisplayName)

		oneWasPartnered = oneWasPartnered || v.HostPartnered
	}
	require.True(t, oneWasPartnered)

	return true
}
