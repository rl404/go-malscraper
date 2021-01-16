package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTopAnime(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetTopAnime(0, 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.Equal(t, len(d), 50)
	emptyEp := true
	emptyY1, emptyM1, emptyY2, emptyM2 := true, true, true, true
	for _, a := range d {
		assert.NotZero(t, a.Rank)
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, a.Image)
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Type)
		if a.Episode > 0 {
			emptyEp = false
		}
		if a.StartDate.Year != 0 {
			emptyY1 = false
		}
		if a.StartDate.Month != 0 {
			emptyM1 = false
		}
		if a.EndDate.Year != 0 {
			emptyY2 = false
		}
		if a.EndDate.Month != 0 {
			emptyM2 = false
		}
		assert.NotZero(t, a.Member)
		assert.NotZero(t, a.Score)
	}
	assert.False(t, emptyEp)
	assert.False(t, emptyY1)
	assert.False(t, emptyM1)
	assert.False(t, emptyY2)
	assert.False(t, emptyM2)
	time.Sleep(sleepDur)
}

func TestGetTopManga(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetTopManga(0, 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.Equal(t, len(d), 50)
	emptyVol := true
	emptyY1, emptyM1, emptyY2, emptyM2 := true, true, true, true
	for _, a := range d {
		assert.NotZero(t, a.Rank)
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, a.Image)
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Type)
		if a.Volume != 0 {
			emptyVol = false
		}
		if a.StartDate.Year != 0 {
			emptyY1 = false
		}
		if a.StartDate.Month != 0 {
			emptyM1 = false
		}
		if a.EndDate.Year != 0 {
			emptyY2 = false
		}
		if a.EndDate.Month != 0 {
			emptyM2 = false
		}
		assert.NotZero(t, a.Member)
		assert.NotZero(t, a.Score)
	}
	assert.False(t, emptyVol)
	assert.False(t, emptyY1)
	assert.False(t, emptyM1)
	assert.False(t, emptyY2)
	assert.False(t, emptyM2)
	time.Sleep(sleepDur)
}

func TestGetTopCharacter(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetTopCharacter(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.Equal(t, len(d), 50)
	emptyJap := true
	for _, a := range d {
		assert.NotZero(t, a.Rank)
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Image)
		if a.JapaneseName != "" {
			emptyJap = false
		}
		assert.NotZero(t, a.Favorite)
	}
	assert.False(t, emptyJap)
	time.Sleep(sleepDur)
}

func TestGetTopPeople(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetTopPeople(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.Equal(t, len(d), 50)
	emptyY, emptyM, emptyD := true, true, true
	for _, a := range d {
		assert.NotZero(t, a.Rank)
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.JapaneseName)
		assert.NotEmpty(t, a.Image)
		assert.NotNil(t, a.Birthday)
		if a.Birthday.Year != 0 {
			emptyY = false
		}
		if a.Birthday.Month != 0 {
			emptyM = false
		}
		if a.Birthday.Day != 0 {
			emptyD = false
		}
		assert.NotZero(t, a.Favorite)
	}
	assert.False(t, emptyY)
	assert.False(t, emptyM)
	assert.False(t, emptyD)
	time.Sleep(sleepDur)
}
