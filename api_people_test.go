package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPeople(t *testing.T) {
	d, code, err := mal.GetPeople(25)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, d.ID)
	assert.NotEmpty(t, d.Name)
	assert.NotEmpty(t, d.Image)
	assert.NotEmpty(t, d.GivenName)
	assert.NotEmpty(t, d.FamilyName)
	assert.NotZero(t, len(d.AlternativeNames))
	for _, a := range d.AlternativeNames {
		assert.NotEmpty(t, a)
	}
	assert.NotZero(t, d.Birthday.Year)
	assert.NotZero(t, d.Birthday.Month)
	assert.NotZero(t, d.Birthday.Day)
	assert.NotEmpty(t, d.Website)
	assert.NotZero(t, d.Favorite)
	assert.NotEmpty(t, d.More)
	time.Sleep(sleepDur)
}

func TestGetPeopleCharacter(t *testing.T) {
	d, code, err := mal.GetPeopleCharacter(14)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyAnimeImg, emptyCharImg := true, true
	for _, c := range d {
		assert.NotZero(t, c.Anime.ID)
		assert.NotEmpty(t, c.Anime.Name)
		assert.NotEmpty(t, c.Anime.Role)
		assert.NotZero(t, c.Character.ID)
		assert.NotEmpty(t, c.Character.Name)
		assert.NotEmpty(t, c.Character.Role)

		if c.Anime.Image != "" {
			emptyAnimeImg = false
		}

		if c.Character.Image != "" {
			emptyCharImg = false
		}
	}
	assert.False(t, emptyAnimeImg)
	assert.False(t, emptyCharImg)
	time.Sleep(sleepDur)
}

func TestGetPeopleStaff(t *testing.T) {
	d, code, err := mal.GetPeopleStaff(14)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Name)
		assert.NotEmpty(t, c.Image)
		assert.NotEmpty(t, c.Role)
	}
	time.Sleep(sleepDur)
}

func TestGetPeopleManga(t *testing.T) {
	d, code, err := mal.GetPeopleManga(14)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Name)
		assert.NotEmpty(t, c.Image)
		assert.NotEmpty(t, c.Role)
	}
	time.Sleep(sleepDur)
}

func TestGetPeopleNews(t *testing.T) {
	d, code, err := mal.GetPeopleNews(11297)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Title)
		assert.NotEmpty(t, c.Image)
		assert.NotEmpty(t, c.Content)
		assert.NotZero(t, c.Date)
		assert.NotEmpty(t, c.Username)
		assert.NotZero(t, c.ForumID)
		assert.NotZero(t, c.Comment)
	}
	time.Sleep(sleepDur)
}

func TestGetPeopleArticle(t *testing.T) {
	d, code, err := mal.GetPeopleArticle(65)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptySpoiler, emptyAd, emptyTag := true, true, true
	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Title)
		assert.NotEmpty(t, c.Image)
		assert.NotEmpty(t, c.Summary)
		assert.NotEmpty(t, c.Username)
		assert.NotZero(t, c.View)

		if c.IsSpoiler {
			emptySpoiler = false
		}
		if c.IsAdvertorial {
			emptyAd = false
		}
		if len(c.Tags) > 0 {
			emptyTag = false
		}
	}
	assert.False(t, emptySpoiler)
	assert.False(t, emptyAd)
	assert.False(t, emptyTag)
	time.Sleep(sleepDur)
}

func TestGetPeoplePicture(t *testing.T) {
	d, code, err := mal.GetPeoplePicture(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, p := range d {
		assert.NotEmpty(t, p)
	}
	time.Sleep(sleepDur)
}
