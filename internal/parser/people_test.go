package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPeople(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetPeople(33)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetPeople(25)
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
	})
	time.Sleep(sleepDur)
}

func TestGetPeopleCharacter(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetPeopleCharacter(33)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetPeopleCharacter(14)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyImg := true
		for _, c := range d {
			assert.NotZero(t, c.Anime.ID)
			assert.NotEmpty(t, c.Anime.Name)
			assert.NotEmpty(t, c.Anime.Image)
			assert.NotEmpty(t, c.Anime.Role)
			assert.NotZero(t, c.Character.ID)
			assert.NotEmpty(t, c.Character.Name)
			assert.NotEmpty(t, c.Character.Role)

			if c.Character.Image != "" {
				emptyImg = false
			}
		}
		assert.False(t, emptyImg)
	})
	time.Sleep(sleepDur)
}

func TestGetPeopleStaff(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetPeopleStaff(33)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetPeopleStaff(14)
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
	})
	time.Sleep(sleepDur)
}

func TestGetPeopleManga(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetPeopleManga(33)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetPeopleManga(14)
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
	})
	time.Sleep(sleepDur)
}

func TestGetPeopleNews(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetPeopleNews(33)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetPeopleNews(11297)
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
	})
	time.Sleep(sleepDur)
}

func TestGetPeopleArticle(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetPeopleArticle(33)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetPeopleArticle(65)
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
	})
	time.Sleep(sleepDur)
}

func TestGetPeoplePicture(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetPeoplePicture(33)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetPeoplePicture(1)
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
