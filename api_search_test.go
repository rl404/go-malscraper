package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchAnime(t *testing.T) {
	d, code, err := mal.SearchAnime("naruto", 0)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyEp, emptyScore, emptyRated, emptyMember := true, true, true, true
	emptyY1, emptyM1, emptyD1, emptyY2, emptyM2, emptyD2 := true, true, true, true, true, true
	for _, a := range d {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, a.Image)
		assert.NotEmpty(t, a.Summary)
		assert.NotEmpty(t, a.Type)
		if a.Episode > 0 {
			emptyEp = false
		}
		if a.Score > 0 {
			emptyScore = false
		}
		if a.StartDate.Year > 0 {
			emptyY1 = false
		}
		if a.StartDate.Month > 0 {
			emptyM1 = false
		}
		if a.StartDate.Day > 0 {
			emptyD1 = false
		}
		if a.EndDate.Year > 0 {
			emptyY2 = false
		}
		if a.EndDate.Month > 0 {
			emptyM2 = false
		}
		if a.EndDate.Day > 0 {
			emptyD2 = false
		}
		if a.Member > 0 {
			emptyMember = false
		}
		if a.Rated != "" {
			emptyRated = false
		}
	}
	assert.False(t, emptyEp)
	assert.False(t, emptyScore)
	assert.False(t, emptyRated)
	assert.False(t, emptyMember)
	assert.False(t, emptyY1)
	assert.False(t, emptyM1)
	assert.False(t, emptyD1)
	assert.False(t, emptyY2)
	assert.False(t, emptyM2)
	assert.False(t, emptyD2)
	time.Sleep(sleepDur)
}

func TestSearchManga(t *testing.T) {
	d, code, err := mal.SearchManga("naruto", 0)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptySum, emptyCh, emptyScore, emptyVol, emptyMember := true, true, true, true, true
	emptyY1, emptyM1, emptyD1, emptyY2, emptyM2, emptyD2 := true, true, true, true, true, true
	for _, a := range d {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, a.Image)
		if a.Summary != "" {
			emptySum = false
		}
		assert.NotEmpty(t, a.Type)
		if a.Chapter > 0 {
			emptyCh = false
		}
		if a.Volume > 0 {
			emptyVol = false
		}
		if a.Score > 0 {
			emptyScore = false
		}
		if a.StartDate.Year > 0 {
			emptyY1 = false
		}
		if a.StartDate.Month > 0 {
			emptyM1 = false
		}
		if a.StartDate.Day > 0 {
			emptyD1 = false
		}
		if a.EndDate.Year > 0 {
			emptyY2 = false
		}
		if a.EndDate.Month > 0 {
			emptyM2 = false
		}
		if a.EndDate.Day > 0 {
			emptyD2 = false
		}
		if a.Member > 0 {
			emptyMember = false
		}
	}
	assert.False(t, emptySum)
	assert.False(t, emptyCh)
	assert.False(t, emptyScore)
	assert.False(t, emptyMember)
	assert.False(t, emptyVol)
	assert.False(t, emptyY1)
	assert.False(t, emptyM1)
	assert.False(t, emptyD1)
	assert.False(t, emptyY2)
	assert.False(t, emptyM2)
	assert.False(t, emptyD2)
	time.Sleep(sleepDur)
}

func TestSearchCharacter(t *testing.T) {
	d, code, err := mal.SearchCharacter("naruto", 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyNick := true, true
	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Name)
		if c.Image != "" {
			emptyImg = false
		}
		if c.Nickname != "" {
			emptyNick = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyNick)
	time.Sleep(sleepDur)
}

func TestSearchPeople(t *testing.T) {
	d, code, err := mal.SearchPeople("kana", 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyNick := true, true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Nickname != "" {
			emptyNick = false
		}
		if p.Image != "" {
			emptyImg = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyNick)
	time.Sleep(sleepDur)
}

func TestSearchClub(t *testing.T) {
	d, code, err := mal.SearchClub("one", 0)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyMem, emptyCre := true, true, true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Image != "" {
			emptyImg = false
		}
		assert.NotEmpty(t, p.Summary)
		if p.Creator != "" {
			emptyCre = false
		}
		if p.Member > 0 {
			emptyMem = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyMem)
	assert.False(t, emptyCre)
	time.Sleep(sleepDur)
}

func TestSearchUser(t *testing.T) {
	d, code, err := mal.SearchUser("rl404", 0)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyOnline := true,true
	for _, p := range d {
		assert.NotEmpty(t, p.Username)
		if p.Image != "" {
			emptyImg = false
		}
		if p.LastOnline != nil {
			emptyOnline = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyOnline)
}
