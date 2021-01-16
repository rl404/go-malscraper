package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClubs(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetClubs(-1)

	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptySum, emptyMem := true, true, true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Image != "" {
			emptyImg = false
		}
		if p.Summary != "" {
			emptySum = false
		}
		assert.NotEmpty(t, p.Creator)
		if p.Member > 0 {
			emptyMem = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptySum)
	assert.False(t, emptyMem)
	time.Sleep(sleepDur)
}

func TestGetClub(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetClub(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetClub(78933)
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
	})
	time.Sleep(sleepDur)
}

func TestGetClubMember(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetClubMember(0, 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetClubMember(1, 1)
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
	})
	time.Sleep(sleepDur)
}

func TestGetClubPicture(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetClubPicture(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetClubPicture(1)
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

func TestGetClubRelated(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetClubRelated(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetClubRelated(1)
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
	})
	time.Sleep(sleepDur)
}
