package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetManga(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetManga(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetManga(42)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotNil(t, d)
		assert.NotZero(t, d.ID)
		assert.NotEmpty(t, d.Image)
		assert.NotEmpty(t, d.Title)
		assert.NotEmpty(t, d.AlternativeTitles.English)
		assert.NotEmpty(t, d.AlternativeTitles.Synonym)
		assert.NotEmpty(t, d.AlternativeTitles.Japanese)
		assert.NotEmpty(t, d.Synopsis)
		assert.NotZero(t, d.Score)
		assert.NotZero(t, d.Voter)
		assert.NotZero(t, d.Rank)
		assert.NotZero(t, d.Popularity)
		assert.NotZero(t, d.Member)
		assert.NotZero(t, d.Favorite)
		assert.NotEmpty(t, d.Type)
		assert.NotZero(t, d.Volume)
		assert.NotZero(t, d.Chapter)
		assert.NotEmpty(t, d.Status)
		assert.NotZero(t, d.PublishingDate.Start.Year)
		assert.NotZero(t, d.PublishingDate.Start.Month)
		assert.NotZero(t, d.PublishingDate.Start.Day)
		assert.NotZero(t, d.PublishingDate.End.Year)
		assert.NotZero(t, d.PublishingDate.End.Month)
		assert.NotZero(t, d.PublishingDate.End.Day)

		for _, p := range d.Authors {
			assert.NotZero(t, p.ID)
			assert.NotEmpty(t, p.Name)
		}

		for _, l := range d.Serializations {
			assert.NotZero(t, l.ID)
			assert.NotEmpty(t, l.Name)
		}

		for _, g := range d.Genres {
			assert.NotZero(t, g.ID)
			assert.NotEmpty(t, g.Name)
		}

		for _, r := range d.Related.Sequel {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}

		// alternative setting, summary, full story, parent story
		// always empty for this manga.

		for _, r := range d.Related.Prequel {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}

		for _, r := range d.Related.AltVersion {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}

		for _, r := range d.Related.SideStory {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}

		for _, r := range d.Related.SpinOff {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}

		for _, r := range d.Related.Adaptation {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}

		for _, r := range d.Related.Character {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}

		for _, r := range d.Related.Other {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Title)
			assert.NotEmpty(t, r.Type)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetMangaReview(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaReview(0, 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaReview(1, 1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
		require.NotZero(t, len(d))

		emptyImg := true
		story0, art0, char0, enj0 := true, true, true, true
		for _, r := range d {
			assert.NotZero(t, r.ID)
			assert.NotEmpty(t, r.Username)
			assert.NotZero(t, r.Source.ID)
			assert.NotEmpty(t, r.Source.Title)
			assert.NotEmpty(t, r.Source.Image)
			assert.Equal(t, r.Source.Type, "manga")
			assert.NotZero(t, r.Helpful)
			assert.NotZero(t, r.Date)
			assert.NotEmpty(t, r.Chapter)
			assert.NotZero(t, r.Score.Overall)
			assert.NotEmpty(t, r.Review)

			if r.Image != "" {
				emptyImg = false
			}
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
		}

		assert.False(t, emptyImg)
		assert.False(t, story0)
		assert.False(t, art0)
		assert.False(t, char0)
		assert.False(t, enj0)
	})
	time.Sleep(sleepDur)
}

func TestGetMangaRecommendation(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaRecommendation(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaRecommendation(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
		require.NotZero(t, len(d))

		for _, r := range d {
			assert.NotZero(t, r.Source.ID)
			assert.NotEmpty(t, r.Source.Title)
			assert.NotEmpty(t, r.Source.Image)
			assert.Equal(t, r.Source.Type, "manga")
			assert.NotZero(t, r.Recommended.ID)
			assert.NotEmpty(t, r.Recommended.Title)
			assert.NotEmpty(t, r.Recommended.Image)
			assert.Equal(t, r.Recommended.Type, "manga")
			assert.NotZero(t, len(r.Users))

			for _, u := range r.Users {
				assert.NotEmpty(t, u.Username)
				assert.NotEmpty(t, u.Content)
			}
		}
	})
	time.Sleep(sleepDur)
}

func TestGetMangaStats(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaStats(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaStats(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, d.Summary.Current)
		assert.NotZero(t, d.Summary.Completed)
		assert.NotZero(t, d.Summary.OnHold)
		assert.NotZero(t, d.Summary.Dropped)
		assert.NotZero(t, d.Summary.Planned)
		assert.NotZero(t, d.Summary.Total)

		assert.NotZero(t, d.Score.Score1.Vote)
		assert.NotZero(t, d.Score.Score1.Percent)
		assert.NotZero(t, d.Score.Score2.Vote)
		assert.NotZero(t, d.Score.Score2.Percent)
		assert.NotZero(t, d.Score.Score3.Vote)
		assert.NotZero(t, d.Score.Score3.Percent)
		assert.NotZero(t, d.Score.Score4.Vote)
		assert.NotZero(t, d.Score.Score4.Percent)
		assert.NotZero(t, d.Score.Score5.Vote)
		assert.NotZero(t, d.Score.Score5.Percent)
		assert.NotZero(t, d.Score.Score6.Vote)
		assert.NotZero(t, d.Score.Score6.Percent)
		assert.NotZero(t, d.Score.Score7.Vote)
		assert.NotZero(t, d.Score.Score7.Percent)
		assert.NotZero(t, d.Score.Score8.Vote)
		assert.NotZero(t, d.Score.Score8.Percent)
		assert.NotZero(t, d.Score.Score9.Vote)
		assert.NotZero(t, d.Score.Score9.Percent)
		assert.NotZero(t, d.Score.Score10.Vote)
		assert.NotZero(t, d.Score.Score10.Percent)
	})
	time.Sleep(sleepDur)
}

func TestGetMangaCharacter(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaCharacter(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaCharacter(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
		require.NotZero(t, len(d))

		for _, c := range d {
			assert.NotZero(t, c.ID)
			assert.NotEmpty(t, c.Name)
			assert.NotEmpty(t, c.Image)
			assert.NotEmpty(t, c.Role)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetMangaNews(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaNews(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaNews(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
		require.NotZero(t, len(d))

		emptyImg := true
		for _, n := range d {
			assert.NotZero(t, n.ID)
			assert.NotEmpty(t, n.Title)
			assert.NotEmpty(t, n.Content)
			assert.NotZero(t, n.Date)
			assert.NotEmpty(t, n.Username)
			assert.NotZero(t, n.ForumID)
			assert.NotZero(t, n.Comment)

			if n.Image != "" {
				emptyImg = false
			}
		}

		assert.False(t, emptyImg)
	})
	time.Sleep(sleepDur)
}

func TestGetMangaArticle(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaArticle(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaArticle(25)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
		require.NotZero(t, len(d))

		// is_advertorial is not tested because there is no manga
		// featured article with complete test case.

		emptySpoiler, emptyTag := true, true
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
			if len(f.Tags) > 0 {
				emptyTag = false
			}
		}

		assert.False(t, emptySpoiler)
		assert.False(t, emptyTag)
	})
	time.Sleep(sleepDur)
}

func TestGetMangaClub(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaClub(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaClub(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
		require.NotZero(t, len(d))

		for _, c := range d {
			assert.NotZero(t, c.ID)
			assert.NotEmpty(t, c.Name)
			assert.NotZero(t, c.Member)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetMangaPicture(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaPicture(0)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaPicture(1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
		require.NotZero(t, len(d))

		for _, p := range d {
			assert.NotEmpty(t, p)
		}
	})
	time.Sleep(sleepDur)
}

func TestGetMangaMoreInfo(t *testing.T) {
	parser := New(true, true, log)
	t.Run("not-found", func(t *testing.T) {
		d, code, err := parser.GetMangaMoreInfo(0)
		assert.Empty(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMangaMoreInfo(2)
		require.NotEmpty(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)
	})
	time.Sleep(sleepDur)
}
