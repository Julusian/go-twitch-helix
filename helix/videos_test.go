package helix

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const videosUserID = 72938233 // dweia131

func checkAllVideosDataIsWellDefined(t *testing.T, data []VideosEntry) {
	for _, v := range data {
		require.NotEqual(t, "", v.ID)
		require.NotEqual(t, 0, v.UserID)
		require.NotEqual(t, "", v.Title)
		require.NotEqual(t, time.Unix(0, 0), v.CreatedAt)
		require.NotEqual(t, time.Unix(0, 0), v.PublishedAt)
		require.NotEqual(t, "", v.Duration)
	}
}
func checkAllVideosDataAdditional(t *testing.T, data []VideosEntry) {
	descCount := 0
	langCount := 0
	viewCount := 0
	thumbCount := 0

	for _, v := range data {
		if v.Description != "" {
			descCount++
		}
		if v.Language != "" {
			langCount++
		}
		if v.ViewCount != 0 {
			viewCount++
		}
		if v.ThumbnailURL != "" {
			thumbCount++
		}
	}

	// require.NotEqual(t, 0, descCount)
	require.NotEqual(t, 0, langCount)
	// require.NotEqual(t, 0, viewCount)
	require.NotEqual(t, 0, thumbCount)
}

func TestGetVideosByGameId(t *testing.T) {
	tc := newTestClient(t)

	res, _, err := GetVideosByGameId(tc, "27471", &VideosParams{})
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllVideosDataIsWellDefined(t, res.Data)
	checkAllVideosDataAdditional(t, res.Data)
}

func TestGetVideosByGameIdEmptyString(t *testing.T) {
	tc := newTestClient(t)

	res, _, err := GetVideosByGameId(tc, "", &VideosParams{})
	require.NotNil(t, err)
	require.Nil(t, res)
}

func TestGetVideosByUserId(t *testing.T) {
	tc := newTestClient(t)

	res, _, err := GetVideosByUserId(tc, videosUserID, &VideosParams{})
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllVideosDataIsWellDefined(t, res.Data)
}
