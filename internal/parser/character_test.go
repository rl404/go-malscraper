package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCharacter(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetCharacter(-1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetCharacter(8)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetCharacter(40)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.ID)
		assert.NotEmpty(t, d.Image)
		assert.NotEmpty(t, d.Nickname)
		assert.NotEmpty(t, d.Name)
		assert.NotEmpty(t, d.JapaneseName)
		assert.NotZero(t, d.Favorite)
		assert.NotEmpty(t, d.About)
	})
	time.Sleep(sleepDur)
}

func TestGetCharacterArticle(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetCharacterArticle(-1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetCharacterArticle(8)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetCharacterArticle(5432)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptySpoiler, emptyAd, emptyTag := true, true, true
		for _, f := range d {
			assert.NotZero(t, f.ID)
			assert.NotEmpty(t, f.Title)
			assert.NotEmpty(t, f.Image)
			assert.NotEmpty(t, f.Summary)
			assert.NotEmpty(t, f.Username)
			assert.NotZero(t, f.View)

			if f.IsSpoiler {
				emptySpoiler = false
			}
			if f.IsAdvertorial {
				emptyAd = false
			}
			if len(f.Tags) > 0 {
				emptyTag = false
			}
		}
		assert.False(t, emptySpoiler)
		assert.False(t, emptyAd)
		assert.False(t, emptyTag)
	})
	time.Sleep(sleepDur)
}

func TestGetCharacterOgraphy(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetCharacterOgraphy("anime", -1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetCharacterOgraphy("anime", 8)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("ok-anime", func(t *testing.T) {
		d, code, err := parser.GetCharacterOgraphy("anime", 1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		for _, a := range d {
			assert.NotZero(t, a.ID)
			assert.NotEmpty(t, a.Name)
			assert.NotEmpty(t, a.Image)
			assert.NotEmpty(t, a.Role)
		}
	})

	t.Run("ok-manga", func(t *testing.T) {
		d, code, err := parser.GetCharacterOgraphy("manga", 1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		for _, a := range d {
			assert.NotZero(t, a.ID)
			assert.NotEmpty(t, a.Name)
			assert.NotEmpty(t, a.Image)
			assert.NotEmpty(t, a.Role)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetCharacterPicture(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetCharacterPicture(-1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetCharacterPicture(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		for _, p := range d {
			assert.NotEmpty(t, p)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetCharacterClub(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetCharacterClub(-1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetCharacterClub(8)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetCharacterClub(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		for _, c := range d {
			assert.NotZero(t, c.ID)
			assert.NotEmpty(t, c.Name)
			assert.NotZero(t, c.Member)
		}
	})
	time.Sleep(sleepDur)
}
