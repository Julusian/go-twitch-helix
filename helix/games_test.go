package helix

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func checkAllGamesDataIsWellDefined(t *testing.T, data []GamesEntry) {
	for _, v := range data {
		require.NotEqual(t, "", v.ID)
		require.NotEqual(t, "", v.Name)
		require.NotEqual(t, "", v.BoxArtURL)
	}
}

func TestGamesById(t *testing.T) {
	tc := newTestClient(t)

	res, _, err := GetGamesById(tc, "", "1234", "345", "99999999")
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllGamesDataIsWellDefined(t, res.Data)

	ids := []string{}
	for _, v := range res.Data {
		ids = append(ids, v.ID)
	}
	sort.Strings(ids)
	require.Equal(t, []string{"1234", "345"}, ids)
}

func TestGamesByName(t *testing.T) {
	tc := newTestClient(t)

	res, _, err := GetGamesByName(tc, "", "Minecraft", "Not a real game I hope!", "Prison Architect")
	require.Nil(t, err)
	require.NotNil(t, res)

	require.NotEqual(t, 0, len(res.Data))
	checkAllGamesDataIsWellDefined(t, res.Data)

	ids := []string{}
	for _, v := range res.Data {
		ids = append(ids, v.ID)
	}
	sort.Strings(ids)
	require.Equal(t, []string{"27471", "32979"}, ids)
}

func TestGamesNoParmas(t *testing.T) {
	tc := newTestClient(t)

	res, _, err := GetGamesById(tc)
	require.NotNil(t, err)
	require.Nil(t, res)
}
