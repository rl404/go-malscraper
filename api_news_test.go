package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNews(t *testing.T) {
	t.Run("related-ok", func(t *testing.T) {
		d, code, err := mal.GetNews(34036779)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.ID)
		assert.NotEmpty(t, d.Title)
		assert.NotEmpty(t, d.Content)
		assert.NotZero(t, d.Date)
		assert.NotEmpty(t, d.Username)
		assert.NotZero(t, d.ForumID)
		assert.NotZero(t, d.Comment)
		assert.NotZero(t, len(d.Related.Anime))
		for _, a := range d.Related.Anime {
			assert.NotZero(t, a.ID)
			assert.NotEmpty(t, a.Name)
		}
		assert.NotZero(t, len(d.Related.Manga))
		for _, m := range d.Related.Manga {
			assert.NotZero(t, m.ID)
			assert.NotEmpty(t, m.Name)
		}
		assert.NotZero(t, len(d.Related.People))
		for _, m := range d.Related.People {
			assert.NotZero(t, m.ID)
			assert.NotEmpty(t, m.Name)
		}
	})

	t.Run("tag-ok", func(t *testing.T) {
		d, code, err := mal.GetNews(61210693)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d.Tags))
		for _, tag := range d.Tags {
			assert.NotEmpty(t, tag)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetNewsList(t *testing.T) {
	t.Run("no-tag-ok", func(t *testing.T) {
		d, code, err := mal.GetNewsList(1, "")
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyCmt := true
		for _, n := range d {
			assert.NotZero(t, n.ID)
			assert.NotEmpty(t, n.Title)
			assert.NotEmpty(t, n.Image)
			assert.NotEmpty(t, n.Content)
			assert.NotZero(t, n.Date)
			assert.NotEmpty(t, n.Username)
			assert.NotZero(t, n.ForumID)
			if n.Comment > 0 {
				emptyCmt = false
			}
		}
		assert.False(t, emptyCmt)
	})

	t.Run("with-tag-ok", func(t *testing.T) {
		d, code, err := mal.GetNewsList(1, "new_anime")
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyCmt := true
		for _, n := range d {
			assert.NotZero(t, n.ID)
			assert.NotEmpty(t, n.Title)
			assert.NotEmpty(t, n.Image)
			assert.NotEmpty(t, n.Content)
			assert.NotZero(t, n.Date)
			assert.NotEmpty(t, n.Username)
			assert.NotZero(t, n.ForumID)
			if n.Comment > 0 {
				emptyCmt = false
			}
		}
		assert.False(t, emptyCmt)
	})
	time.Sleep(sleepDur)
}

func TestGetNewsTag(t *testing.T) {
	d, code, err := mal.GetNewsTag()
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d.Anime))
	for _, a := range d.Anime {
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Tag)
		assert.NotEmpty(t, a.Description)
	}
	assert.NotZero(t, len(d.Manga))
	for _, a := range d.Manga {
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Tag)
		assert.NotEmpty(t, a.Description)
	}
	assert.NotZero(t, len(d.People))
	for _, a := range d.People {
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Tag)
		assert.NotEmpty(t, a.Description)
	}
	assert.NotZero(t, len(d.Music))
	for _, a := range d.Music {
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Tag)
		assert.NotEmpty(t, a.Description)
	}
	assert.NotZero(t, len(d.Event))
	for _, a := range d.Event {
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Tag)
		assert.NotEmpty(t, a.Description)
	}
	assert.NotZero(t, len(d.Industry))
	emptyDesc := true
	for _, a := range d.Industry {
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Tag)
		if a.Description != "" {
			emptyDesc = false
		}
	}
	assert.False(t, emptyDesc)
	time.Sleep(sleepDur)
}
