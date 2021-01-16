package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRecommendation(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetRecommendation("anime", 1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("anime-ok", func(t *testing.T) {
		d, code, err := parser.GetRecommendation("anime", 1, 6)
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
		d, code, err := parser.GetRecommendation("manga", 1, 21)
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
	parser := New(true, true, log)
	t.Run("invalid-type", func(t *testing.T) {
		d, code, err := parser.GetRecommendations("random", 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok-anime", func(t *testing.T) {
		d, code, err := parser.GetRecommendations("anime", 1)
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
		d, code, err := parser.GetRecommendations("manga", 1)
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
