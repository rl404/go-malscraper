package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClubs(t *testing.T) {
	d, code, err := mal.GetClubs(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptySum, emptyMem, emptyCreator := true, true, true, true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Image != "" {
			emptyImg = false
		}
		if p.Summary != "" {
			emptySum = false
		}
		if p.Creator != "" {
			emptyCreator = false
		}
		if p.Member > 0 {
			emptyMem = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptySum)
	assert.False(t, emptyMem)
	assert.False(t, emptyCreator)
	time.Sleep(sleepDur)
}

func TestGetClub(t *testing.T) {
	d, code, err := mal.GetClub(78933)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, d.ID)
	assert.NotEmpty(t, d.Name)
	assert.NotEmpty(t, d.Image)
	assert.NotEmpty(t, d.Information)
	assert.NotEmpty(t, d.Category)
	assert.NotEmpty(t, d.Type)
	assert.NotZero(t, d.Member)
	assert.NotZero(t, d.Picture)
	assert.NotZero(t, d.CreatedDate)
	assert.NotZero(t, len(d.Admins))
	for _, a := range d.Admins {
		assert.NotEmpty(t, a.Username)
		assert.NotZero(t, len(a.Roles))
		for _, r := range a.Roles {
			assert.NotEmpty(t, r)
		}
	}
	time.Sleep(sleepDur)
}

func TestGetClubMember(t *testing.T) {
	d, code, err := mal.GetClubMember(1, 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg := true
	for _, u := range d {
		assert.NotEmpty(t, u.Username)
		if u.Image != "" {
			emptyImg = false
		}
	}
	assert.False(t, emptyImg)
	time.Sleep(sleepDur)
}

func TestGetClubPicture(t *testing.T) {
	d, code, err := mal.GetClubPicture(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, p := range d {
		assert.NotEmpty(t, p)
	}
	time.Sleep(sleepDur)
}

func TestGetClubRelated(t *testing.T) {
	d, code, err := mal.GetClubRelated(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d.Anime))
	for _, a := range d.Anime {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
	}
	assert.NotZero(t, len(d.Manga))
	for _, a := range d.Manga {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
	}
	assert.NotZero(t, len(d.Character))
	for _, a := range d.Character {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
	}
	time.Sleep(sleepDur)
}
