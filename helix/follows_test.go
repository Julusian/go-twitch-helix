package helix

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const followersTestChannel = 79110861 // Julusian
const followersTestUser = 92016198    // Botofdork

func checkAllDataIsWellDefined(t *testing.T, data []FollowsEntry) {
	for _, v := range data {
		require.NotEqual(t, 0, v.ChannelID)
		require.NotEqual(t, 0, v.UserID)
		require.NotEqual(t, time.Unix(0, 0), v.FollowedAt)
	}
}

func TestFollowsForChannel(t *testing.T) {
	tc := newTestClient(t)

	params := &FollowsParams{
		ChannelID: followersTestChannel,
	}

	res, rate, err := GetFollows(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllDataIsWellDefined(t, res.Data)
}

func TestFollowsForChannelPagination(t *testing.T) {
	tc := newTestClient(t)

	params := &FollowsParams{
		ChannelID: followersTestChannel,
		Limit:     1,
	}

	// Get page 1
	res, _, err := GetFollows(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllDataIsWellDefined(t, res.Data)
	firstID := res.Data[0].UserID

	// Get page 2
	params.After = res.Pagination.Cursor
	require.NotEqual(t, "", params.After)

	res, _, err = GetFollows(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllDataIsWellDefined(t, res.Data)
	secondID := res.Data[0].UserID

	require.NotEqual(t, firstID, secondID)

	// Note: Currently returns a 400 error Invalid cursor.
	// // Back to page 1
	// params.Before = res.Pagination.Cursor
	// params.After = ""
	// require.NotEqual(t, "", params.Before)

	// res, _, err = GetFollows(tc, params)
	// require.Nil(t, err)
	// require.NotNil(t, res)

	// require.NotEqual(t, 0, len(res.Data))
	// checkAllDataIsWellDefined(t, res.Data)
	// thirdID := res.Data[0].UserID

	// require.NotEqual(t, thirdID, secondID)
	// require.Equal(t, thirdID, firstID)
}

func TestFollowsForUser(t *testing.T) {
	tc := newTestClient(t)

	params := &FollowsParams{
		UserID: followersTestChannel,
	}

	res, rate, err := GetFollows(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllDataIsWellDefined(t, res.Data)
}

func TestFollowsForUserAndChannel(t *testing.T) {
	tc := newTestClient(t)

	params := &FollowsParams{
		ChannelID: followersTestChannel,
		UserID:    followersTestUser,
	}

	res, rate, err := GetFollows(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllDataIsWellDefined(t, res.Data)
}

func TestFollowsForUserAndChannelNoFollow(t *testing.T) {
	tc := newTestClient(t)

	params := &FollowsParams{
		UserID:    followersTestChannel,
		ChannelID: followersTestUser,
	}

	res, rate, err := GetFollows(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.Equal(t, 0, len(res.Data))
}

func TestFollowsForBadChannel(t *testing.T) {
	tc := newTestClient(t)

	params := &FollowsParams{
		ChannelID: 9,
	}

	res, rate, err := GetFollows(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.Equal(t, 0, len(res.Data))
}
