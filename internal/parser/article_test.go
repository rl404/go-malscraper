package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetArticle(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetArticle(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("related-ok", func(t *testing.T) {
		d, code, err := parser.GetArticle(2321)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.ID)
		assert.NotEmpty(t, d.Title)
		assert.NotEmpty(t, d.Summary)
		assert.NotEmpty(t, d.Content)
		assert.NotZero(t, d.Date)
		assert.NotEmpty(t, d.Username)
		assert.NotZero(t, d.View)
		assert.True(t, d.IsSpoiler)
		assert.True(t, d.IsAdvertorial)
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
		assert.NotZero(t, len(d.Related.Character))
		for _, m := range d.Related.Character {
			assert.NotZero(t, m.ID)
			assert.NotEmpty(t, m.Name)
		}
	})

	t.Run("tag-ok", func(t *testing.T) {
		d, code, err := parser.GetArticle(1821)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d.Tags))
		for _, a := range d.Tags {
			assert.NotEmpty(t, a)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetArticles(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-tag", func(t *testing.T) {
		d, code, err := parser.GetArticles(1, "asd")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("no-tag-ok", func(t *testing.T) {
		d, code, err := parser.GetArticles(1, "")
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptySpo, emptyAd := true, true
		for _, f := range d {
			assert.NotZero(t, f.ID)
			assert.NotEmpty(t, f.Title)
			assert.NotEmpty(t, f.Image)
			assert.NotEmpty(t, f.Summary)
			assert.NotEmpty(t, f.Username)
			assert.NotZero(t, f.View)
			if f.IsAdvertorial {
				emptyAd = false
			}
			if f.IsSpoiler {
				emptySpo = false
			}
		}
		assert.False(t, emptyAd)
		assert.False(t, emptySpo)
	})

	t.Run("with-tag-ok", func(t *testing.T) {
		d, code, err := parser.GetArticles(1, "recap")
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		for _, f := range d {
			assert.NotZero(t, len(f.Tags))
			for _, a := range f.Tags {
				assert.NotEmpty(t, a)
			}
		}
	})
	time.Sleep(sleepDur)
}

func TestGetArticleTag(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetArticleTag()
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, tag := range d {
		assert.NotEmpty(t, tag.Name)
		assert.NotEmpty(t, tag.Tag)
	}
	time.Sleep(sleepDur)
}
