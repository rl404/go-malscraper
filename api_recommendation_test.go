package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRecommendation(t *testing.T) {
	t.Run("anime-ok", func(t *testing.T) {
		d, code, err := mal.GetRecommendationAnime(1, 6)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.Source.ID)
		assert.NotEmpty(t, d.Source.Title)
		assert.NotEmpty(t, d.Source.Image)
		assert.Equal(t, d.Source.Type, "anime")
		assert.NotZero(t, d.Recommended.ID)
		assert.NotEmpty(t, d.Recommended.Title)
		assert.NotEmpty(t, d.Recommended.Image)
		assert.Equal(t, d.Recommended.Type, "anime")
		assert.NotZero(t, len(d.Users))
		for _, u := range d.Users {
			assert.NotEmpty(t, u.Username)
			assert.NotEmpty(t, u.Content)
		}
	})

	t.Run("manga-ok", func(t *testing.T) {
		d, code, err := mal.GetRecommendationManga(1, 21)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.Source.ID)
		assert.NotEmpty(t, d.Source.Title)
		assert.NotEmpty(t, d.Source.Image)
		assert.Equal(t, d.Source.Type, "manga")
		assert.NotZero(t, d.Recommended.ID)
		assert.NotEmpty(t, d.Recommended.Title)
		assert.NotEmpty(t, d.Recommended.Image)
		assert.Equal(t, d.Recommended.Type, "manga")
		assert.NotZero(t, len(d.Users))
		for _, u := range d.Users {
			assert.NotEmpty(t, u.Username)
			assert.NotEmpty(t, u.Content)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetRecommendations(t *testing.T) {
	t.Run("ok-anime", func(t *testing.T) {
		d, code, err := mal.GetAnimeRecommendations(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyImg1, emptyImg2 := true, true
		for _, r := range d {
			assert.NotZero(t, r.Source.ID)
			assert.NotEmpty(t, r.Source.Title)
			if r.Source.Image != "" {
				emptyImg1 = false
			}
			assert.Equal(t, r.Source.Type, "anime")
			assert.NotZero(t, r.Recommended.ID)
			assert.NotEmpty(t, r.Recommended.Title)
			if r.Recommended.Image != "" {
				emptyImg2 = false
			}
			assert.Equal(t, r.Recommended.Type, "anime")
			assert.Equal(t, len(r.Users), 1)
			for _, u := range r.Users {
				assert.NotEmpty(t, u.Username)
				assert.NotEmpty(t, u.Content)
			}
		}
		assert.False(t, emptyImg1)
		assert.False(t, emptyImg2)
	})

	t.Run("ok-manga", func(t *testing.T) {
		d, code, err := mal.GetMangaRecommendations(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyImg1, emptyImg2 := true, true
		for _, r := range d {
			assert.NotZero(t, r.Source.ID)
			assert.NotEmpty(t, r.Source.Title)
			if r.Source.Image != "" {
				emptyImg1 = false
			}
			assert.Equal(t, r.Source.Type, "manga")
			assert.NotZero(t, r.Recommended.ID)
			assert.NotEmpty(t, r.Recommended.Title)
			if r.Recommended.Image != "" {
				emptyImg2 = false
			}
			assert.Equal(t, r.Recommended.Type, "manga")
			assert.Equal(t, len(r.Users), 1)
			for _, u := range r.Users {
				assert.NotEmpty(t, u.Username)
				assert.NotEmpty(t, u.Content)
			}
		}
		assert.False(t, emptyImg1)
		assert.False(t, emptyImg2)
	})
	time.Sleep(sleepDur)
}
