package helix

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func checkAllStreamDataIsWellDefined(t *testing.T, data []StreamsEntry) {
	for _, v := range data {
		require.NotEqual(t, 0, v.ID)
		require.NotEqual(t, "", v.Language)
		require.NotEqual(t, time.Unix(0, 0), v.StartedAt)
		require.NotEqual(t, "", v.ThumbnailURL)
		require.NotEqual(t, "", v.Title) // TODO - this may need moving to a seperate loop as it can sometimes be legitimately empty
		require.NotEqual(t, "", v.Type)
		require.NotEqual(t, 0, v.UserID)
	}
}

func checkStreamsHasValidCommunityID(t *testing.T, data []StreamsEntry) {
	for _, v := range data {
		if len(v.CommunityIDs) > 0 {
			for _, i := range v.CommunityIDs {
				if len(i) > 0 {
					return
				}
			}
		}
	}

	require.Fail(t, "No community_ids were defined in any stream")
}
func checkStreamsHasValidGameID(t *testing.T, data []StreamsEntry) {
	for _, v := range data {
		if v.GameID != "" {
			return
		}
	}

	require.Fail(t, "No community_ids were defined in any stream")
}
func checkStreamsHasValidViewerCount(t *testing.T, data []StreamsEntry) {
	for _, v := range data {
		if v.ViewerCount > 0 {
			return
		}
	}

	require.Fail(t, "No community_ids were defined in any stream")
}

// TODO - Before, After, UserLogin

func TestStreamsAll(t *testing.T) {
	tc := newTestClient(t)

	params := &StreamsParams{
		Limit: 20, // Ensure we get a good sample of data for the full json tests
	}

	res, rate, err := GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)
	checkStreamsHasValidCommunityID(t, res.Data)
	checkStreamsHasValidGameID(t, res.Data)
	checkStreamsHasValidViewerCount(t, res.Data)
}

func TestStreamsByUser(t *testing.T) {
	tc := newTestClient(t)

	params := &StreamsParams{
		Limit: 5,
	}

	res, rate, err := GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.Equal(t, 5, len(res.Data))
	userIds := []int{res.Data[1].UserID, res.Data[3].UserID}

	params.UserID = userIds

	res, _, err = GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, 2, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)

	require.Equal(t, userIds[0], res.Data[0].UserID)
	require.Equal(t, userIds[1], res.Data[1].UserID)
}

func TestStreamsByCommunity(t *testing.T) {
	tc := newTestClient(t)

	params := &StreamsParams{
		Limit: 20,
	}

	res, rate, err := GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	communityID := ""
	for _, v := range res.Data {
		if len(v.CommunityIDs) > 0 {
			communityID = v.CommunityIDs[0]
			break
		}
	}
	require.NotEqual(t, "", communityID)

	params.CommunityID = []string{communityID}

	res, _, err = GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)

	for _, v := range res.Data {
		require.True(t, inArrayStr(communityID, v.CommunityIDs))
	}
}

func inArrayStr(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func TestStreamsByLanguage(t *testing.T) {
	tc := newTestClient(t)

	params := &StreamsParams{
		Limit:    20,
		Language: []string{"en"},
	}

	res, _, err := GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)

	for _, v := range res.Data {
		require.Equal(t, "en", v.Language)
	}
}

func TestStreamsByType(t *testing.T) {
	tc := newTestClient(t)

	params := &StreamsParams{
		Limit: 10,
		Type:  "vodcast",
	}

	// Vodcast
	res, _, err := GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)

	for _, v := range res.Data {
		require.Equal(t, params.Type, v.Type)
	}

	// Live
	params.Type = "live"
	res, _, err = GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)

	for _, v := range res.Data {
		require.Equal(t, params.Type, v.Type)
	}

	// All
	params.Type = "all"
	res, _, err = GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)
}

func TestStreamsByGame(t *testing.T) {
	tc := newTestClient(t)

	params := &StreamsParams{
		Limit: 2,
	}

	res, rate, err := GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, rate)
	require.NotNil(t, res)

	require.Equal(t, 2, len(res.Data))
	gameID := res.Data[1].GameID

	params.GameID = []string{gameID}
	params.Limit = 0

	res, _, err = GetStreams(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllStreamDataIsWellDefined(t, res.Data)

	for _, v := range res.Data {
		require.Equal(t, gameID, v.GameID)
	}
}
