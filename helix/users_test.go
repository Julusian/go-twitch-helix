package helix

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

const usersTest1 = 79110861 // Julusian
const usersTest1Name = "Julusian"
const usersTest2 = 92016198 // Botofdork
const usersTest2Name = "BotOfDork"
const usersTestFake = 9
const usersTestFakeName = "NotARealUserIReallyHope"

func checkAllUserDataIsWellDefined(t *testing.T, data []UsersEntry) {
	for _, v := range data {
		require.NotEqual(t, 0, v.ID)
		require.NotEqual(t, "", v.DisplayName)
		require.NotEqual(t, "", v.Login)
		require.NotEqual(t, "", v.ProfileImageURL)
		require.NotEqual(t, 0, v.ViewCount)
	}
}

func TestUserSelf(t *testing.T) {
	tc := newTestClient(t)

	params := &UsersParams{
		AuthToken: os.Getenv("TWITCH_AUTHTOKEN"),
	}

	res, _, err := GetUsers(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, 1, len(res.Data))
	checkAllUserDataIsWellDefined(t, res.Data)
	require.NotEqual(t, "", res.Data[0].Email)
	require.NotEqual(t, "", res.Data[0].Description)
	require.NotEqual(t, "", res.Data[0].OfflineImageURL)
}

func TestUserById(t *testing.T) {
	tc := newTestClient(t)

	params := &UsersParams{
		IDs: []int{usersTest1, usersTest2, usersTestFake},
	}

	res, _, err := GetUsers(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, 2, len(res.Data))
	checkAllUserDataIsWellDefined(t, res.Data)

	names := make([]string, 2)
	for i, v := range res.Data {
		names[i] = v.DisplayName
	}
	sort.Strings(names)

	expected := []string{usersTest1Name, usersTest2Name}
	sort.Strings(expected)

	require.Equal(t, expected, names)
}

func TestUserByName(t *testing.T) {
	tc := newTestClient(t)

	params := &UsersParams{
		Logins: []string{usersTest1Name, usersTest2Name, usersTestFakeName},
	}

	res, _, err := GetUsers(tc, params)
	require.Nil(t, err)
	require.NotNil(t, res)

	require.Equal(t, 2, len(res.Data))
	checkAllUserDataIsWellDefined(t, res.Data)

	ids := make([]int, 2)
	for i, v := range res.Data {
		ids[i] = v.ID
	}
	sort.Ints(ids)

	expected := []int{usersTest1, usersTest2}
	sort.Ints(expected)

	require.Equal(t, expected, ids)
}
