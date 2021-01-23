package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAnime(t *testing.T) {
	d, code, err := mal.GetAnime(1412)
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
	assert.NotEmpty(t, d.Video)
	assert.NotEmpty(t, d.Synopsis)
	assert.NotZero(t, d.Score)
	assert.NotZero(t, d.Voter)
	assert.NotZero(t, d.Rank)
	assert.NotZero(t, d.Popularity)
	assert.NotZero(t, d.Member)
	assert.NotEmpty(t, d.Favorite)
	assert.NotEmpty(t, d.Type)
	assert.NotZero(t, d.Episode)
	assert.NotEmpty(t, d.Status)
	assert.NotZero(t, d.AiringDate.Start.Year)
	assert.NotZero(t, d.AiringDate.Start.Month)
	assert.NotZero(t, d.AiringDate.Start.Day)
	assert.NotZero(t, d.AiringDate.End.Year)
	assert.NotZero(t, d.AiringDate.End.Month)
	assert.NotZero(t, d.AiringDate.End.Day)
	assert.NotEmpty(t, d.Premiered)
	assert.NotEmpty(t, d.Broadcast)
	assert.NotEmpty(t, d.Source)
	assert.NotEmpty(t, d.Duration)
	assert.NotEmpty(t, d.Rating)

	for _, p := range d.Producers {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
	}

	for _, l := range d.Licensors {
		assert.NotZero(t, l.ID)
		assert.NotEmpty(t, l.Name)
	}

	for _, s := range d.Studios {
		assert.NotZero(t, s.ID)
		assert.NotEmpty(t, s.Name)
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
	// always empty for this anime.

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

	for _, song := range d.Song.Opening {
		assert.NotEmpty(t, song)
	}
	time.Sleep(sleepDur)
}

func TestGetAnimeCharacter(t *testing.T) {
	d, code, err := mal.GetAnimeCharacter(32281)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Name)
		assert.NotEmpty(t, c.Image)
		assert.NotEmpty(t, c.Role)

		emptyImg := true
		for _, v := range c.VoiceActors {
			assert.NotZero(t, v.ID)
			assert.NotEmpty(t, v.Name)
			assert.NotEmpty(t, v.Role)

			if v.Image != "" {
				emptyImg = false
			}
		}
		assert.False(t, emptyImg)
	}
	time.Sleep(sleepDur)
}

func TestGetAnimeVideo(t *testing.T) {
	d, code, err := mal.GetAnimeVideo(1, 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d.Episodes))
	assert.NotZero(t, len(d.Promotions))

	for _, e := range d.Episodes {
		assert.NotZero(t, e.Episode)
		assert.NotEmpty(t, e.Title)
		assert.NotEmpty(t, e.Link)
	}

	for _, p := range d.Promotions {
		assert.NotEmpty(t, p.Title)
		assert.NotEmpty(t, p.Link)
	}
	time.Sleep(sleepDur)
}

func TestGetAnimeEpisode(t *testing.T) {
	d, code, err := mal.GetAnimeEpisode(20, 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	tagEmpty := true
	for _, e := range d {
		assert.NotZero(t, e.Episode)
		assert.NotEmpty(t, e.Title)
		assert.NotEmpty(t, e.JapaneseTitle)
		assert.NotEmpty(t, e.AiredDate)
		assert.NotEmpty(t, e.Link)
		if e.Tag != "" {
			tagEmpty = false
		}
	}
	assert.False(t, tagEmpty)
	time.Sleep(sleepDur)
}

func TestGetAnimeStats(t *testing.T) {
	d, code, err := mal.GetAnimeStats(1)
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
	time.Sleep(sleepDur)
}

func TestGetAnimeReview(t *testing.T) {
	d, code, err := mal.GetAnimeReview(1, 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	emptyImg := true
	story0, anim0, sou0, char0, enj0 := true, true, true, true, true
	for _, r := range d {
		assert.NotZero(t, r.ID)
		assert.NotEmpty(t, r.Username)
		assert.NotZero(t, r.Source.ID)
		assert.NotEmpty(t, r.Source.Title)
		assert.NotEmpty(t, r.Source.Image)
		assert.Equal(t, r.Source.Type, "anime")
		assert.NotZero(t, r.Helpful)
		assert.NotZero(t, r.Date)
		assert.NotEmpty(t, r.Episode)
		assert.NotZero(t, r.Score.Overall)
		assert.NotEmpty(t, r.Review)

		if r.Image != "" {
			emptyImg = false
		}
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
	}

	assert.False(t, emptyImg)
	assert.False(t, story0)
	assert.False(t, anim0)
	assert.False(t, sou0)
	assert.False(t, char0)
	assert.False(t, enj0)
	time.Sleep(sleepDur)
}

func TestGetAnimeRecommendation(t *testing.T) {
	d, code, err := mal.GetAnimeRecommendation(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	for _, r := range d {
		assert.NotZero(t, r.Source.ID)
		assert.NotEmpty(t, r.Source.Title)
		assert.NotEmpty(t, r.Source.Image)
		assert.Equal(t, r.Source.Type, "anime")
		assert.NotZero(t, r.Recommended.ID)
		assert.NotEmpty(t, r.Recommended.Title)
		assert.NotEmpty(t, r.Recommended.Image)
		assert.Equal(t, r.Recommended.Type, "anime")
		assert.NotZero(t, len(r.Users))

		for _, u := range r.Users {
			assert.NotEmpty(t, u.Username)
			assert.NotEmpty(t, u.Content)
		}
	}
	time.Sleep(sleepDur)
}

func TestGetAnimeStaff(t *testing.T) {
	d, code, err := mal.GetAnimeStaff(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	emptyImg := true
	for _, s := range d {
		assert.NotZero(t, s.ID)
		assert.NotEmpty(t, s.Name)
		assert.NotEmpty(t, s.Role)

		if s.Image != "" {
			emptyImg = false
		}
	}
	assert.False(t, emptyImg)
	time.Sleep(sleepDur)
}

func TestGetAnimeNews(t *testing.T) {
	d, code, err := mal.GetAnimeNews(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	emptyComment := true
	for _, n := range d {
		assert.NotZero(t, n.ID)
		assert.NotEmpty(t, n.Title)
		assert.NotEmpty(t, n.Image)
		assert.NotEmpty(t, n.Content)
		assert.NotZero(t, n.Date)
		assert.NotEmpty(t, n.Username)
		assert.NotZero(t, n.ForumID)

		if n.Comment != 0 {
			emptyComment = false
		}
	}
	assert.False(t, emptyComment)
	time.Sleep(sleepDur)
}

func TestGetAnimeArticle(t *testing.T) {
	d, code, err := mal.GetAnimeArticle(235)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

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
	time.Sleep(sleepDur)
}

func TestGetAnimeClub(t *testing.T) {
	d, code, err := mal.GetAnimeClub(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Name)
		assert.NotZero(t, c.Member)
	}
	time.Sleep(sleepDur)
}

func TestGetAnimePicture(t *testing.T) {
	d, code, err := mal.GetAnimePicture(1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	require.NotZero(t, len(d))

	for _, p := range d {
		assert.NotEmpty(t, p)
	}
	time.Sleep(sleepDur)
}

func TestGetAnimeMoreInfo(t *testing.T) {
	d, code, err := mal.GetAnimeMoreInfo(1)
	require.NotEmpty(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)
	time.Sleep(sleepDur)
}
