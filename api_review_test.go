package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetReview(t *testing.T) {
	t.Run("anime-ok", func(t *testing.T) {
		d, code, err := mal.GetReview(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.ID)
		assert.NotEmpty(t, d.Username)
		assert.NotEmpty(t, d.Image)
		assert.NotZero(t, d.Source.ID)
		assert.NotEmpty(t, d.Source.Title)
		assert.NotEmpty(t, d.Source.Image)
		assert.Equal(t, d.Source.Type, "anime")
		assert.NotZero(t, d.Helpful)
		assert.NotZero(t, d.Date)
		assert.NotEmpty(t, d.Episode)
		assert.NotZero(t, d.Score.Overall)
		assert.NotZero(t, d.Score.Story)
		assert.NotZero(t, d.Score.Art)
		assert.NotZero(t, d.Score.Sound)
		assert.NotZero(t, d.Score.Character)
		assert.NotZero(t, d.Score.Enjoyment)
		assert.NotEmpty(t, d.Review)
	})

	t.Run("manga-ok", func(t *testing.T) {
		d, code, err := mal.GetReview(209071)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.ID)
		assert.NotEmpty(t, d.Username)
		assert.NotEmpty(t, d.Image)
		assert.NotZero(t, d.Source.ID)
		assert.NotEmpty(t, d.Source.Title)
		assert.NotEmpty(t, d.Source.Image)
		assert.Equal(t, d.Source.Type, "manga")
		assert.NotZero(t, d.Helpful)
		assert.NotZero(t, d.Date)
		assert.NotEmpty(t, d.Chapter)
		assert.NotZero(t, d.Score.Overall)
		assert.NotZero(t, d.Score.Story)
		assert.NotZero(t, d.Score.Art)
		assert.NotZero(t, d.Score.Character)
		assert.NotZero(t, d.Score.Enjoyment)
		assert.NotEmpty(t, d.Review)
	})
	time.Sleep(sleepDur)
}

func TestGetReviews(t *testing.T) {
	t.Run("ok-anime", func(t *testing.T) {
		d, code, err := mal.GetAnimeReviews(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyImg1, emptyImg2, emptyHelp := true, true, true
		story0, anim0, sou0, char0, enj0 := true, true, true, true, true
		for _, r := range d {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Username)
			if r.Image != "" {
				emptyImg1 = false
			}
			assert.NotZero(t, r.Source.ID)
			assert.NotEmpty(t, r.Source.Title)
			if r.Source.Image != "" {
				emptyImg2 = false
			}
			assert.Equal(t, r.Source.Type, "anime")
			if r.Helpful > 0 {
				emptyHelp = false
			}
			assert.NotZero(t, r.Date)
			assert.NotEmpty(t, r.Episode)
			assert.NotZero(t, r.Score.Overall)
			if r.Score.Story != 0 {
				story0 = false
			}
			if r.Score.Art != 0 {
				anim0 = false
			}
			if r.Score.Sound != 0 {
				sou0 = false
			}
			if r.Score.Character != 0 {
				char0 = false
			}
			if r.Score.Enjoyment != 0 {
				enj0 = false
			}
			assert.NotEmpty(t, r.Review)
		}
		assert.False(t, emptyImg1)
		assert.False(t, emptyImg2)
		assert.False(t, emptyHelp)
		assert.False(t, story0)
		assert.False(t, anim0)
		assert.False(t, sou0)
		assert.False(t, char0)
		assert.False(t, enj0)
	})

	t.Run("ok-manga", func(t *testing.T) {
		d, code, err := mal.GetMangaReviews(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyImg1, emptyImg2, emptyHelp := true, true, true
		story0, art0, char0, enj0 := true, true, true, true
		for _, r := range d {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Username)
			if r.Image != "" {
				emptyImg1 = false
			}
			assert.NotZero(t, r.Source.ID)
			assert.NotEmpty(t, r.Source.Title)
			if r.Source.Image != "" {
				emptyImg2 = false
			}
			assert.Equal(t, r.Source.Type, "manga")
			if r.Helpful > 0 {
				emptyHelp = false
			}
			assert.NotZero(t, r.Date)
			assert.NotEmpty(t, r.Chapter)
			assert.NotZero(t, r.Score.Overall)
			if r.Score.Story != 0 {
				story0 = false
			}
			if r.Score.Art != 0 {
				art0 = false
			}
			if r.Score.Character != 0 {
				char0 = false
			}
			if r.Score.Enjoyment != 0 {
				enj0 = false
			}
			assert.NotEmpty(t, r.Review)
		}
		assert.False(t, emptyImg1)
		assert.False(t, emptyImg2)
		assert.False(t, emptyHelp)
		assert.False(t, story0)
		assert.False(t, art0)
		assert.False(t, char0)
		assert.False(t, enj0)
	})

	t.Run("ok-best", func(t *testing.T) {
		d, code, err := mal.GetBestReviews(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		story0, art0, sou0, char0, enj0 := true, true, true, true, true
		for _, r := range d {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Username)
			assert.NotEmpty(t, r.Image)
			assert.NotZero(t, r.Source.ID)
			assert.NotEmpty(t, r.Source.Title)
			assert.NotEmpty(t, r.Source.Image)
			assert.NotEmpty(t, r.Source.Type)
			assert.NotEmpty(t, r.Helpful)
			assert.NotZero(t, r.Date)
			if r.Source.Type == "anime" {
				assert.NotEmpty(t, r.Episode)
			} else if r.Source.Type == "manga" {
				assert.NotEmpty(t, r.Chapter)
			}
			assert.NotZero(t, r.Score.Overall)
			if r.Score.Story != 0 {
				story0 = false
			}
			if r.Score.Art != 0 {
				art0 = false
			}
			if r.Source.Type == "anime" && r.Score.Sound != 0 {
				sou0 = false
			}
			if r.Score.Character != 0 {
				char0 = false
			}
			if r.Score.Enjoyment != 0 {
				enj0 = false
			}
			assert.NotEmpty(t, r.Review)
		}
		assert.False(t, story0)
		assert.False(t, art0)
		assert.False(t, sou0)
		assert.False(t, char0)
		assert.False(t, enj0)
	})
	time.Sleep(sleepDur)
}
