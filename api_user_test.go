package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	d, code, err := mal.GetUser("Archaeon")
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotEmpty(t, d.Username)
	assert.NotEmpty(t, d.Image)
	assert.NotZero(t, d.LastOnline)
	assert.NotEmpty(t, d.Gender)
	assert.NotZero(t, d.Birthday.Year)
	assert.NotZero(t, d.Birthday.Month)
	assert.NotZero(t, d.Birthday.Day)
	assert.NotEmpty(t, d.Location)
	assert.NotZero(t, d.JoinedDate)
	assert.NotZero(t, d.ForumPost)
	assert.NotZero(t, d.Review)
	assert.NotZero(t, d.Recommendation)
	assert.NotZero(t, d.BlogPost)
	assert.NotZero(t, d.Club)
	assert.NotZero(t, d.Friend)
	assert.NotZero(t, len(d.Sns))
	assert.NotEmpty(t, d.About)
	time.Sleep(sleepDur)
}

func TestGetUserStats(t *testing.T) {
	d, code, err := mal.GetUserStats("kuroikikyou")
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, d.Anime.Days)
	assert.NotZero(t, d.Anime.MeanScore)
	assert.NotZero(t, d.Anime.Current)
	assert.NotZero(t, d.Anime.Completed)
	assert.NotZero(t, d.Anime.OnHold)
	assert.NotZero(t, d.Anime.Dropped)
	assert.NotZero(t, d.Anime.Planned)
	assert.NotZero(t, d.Anime.Total)
	assert.NotZero(t, d.Anime.Rewatched)
	assert.NotZero(t, d.Anime.Episode)
	assert.NotZero(t, d.Manga.Days)
	assert.NotZero(t, d.Manga.MeanScore)
	assert.NotZero(t, d.Manga.Current)
	assert.NotZero(t, d.Manga.Completed)
	assert.NotZero(t, d.Manga.OnHold)
	assert.NotZero(t, d.Manga.Dropped)
	assert.NotZero(t, d.Manga.Planned)
	assert.NotZero(t, d.Manga.Total)
	assert.NotZero(t, d.Manga.Reread)
	assert.NotZero(t, d.Manga.Chapter)
	assert.NotZero(t, d.Manga.Volume)
	time.Sleep(sleepDur)
}

func TestGetUserFavorite(t *testing.T) {
	d, code, err := mal.GetUserFavorite("rl404")
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d.Anime))
	for _, a := range d.Anime {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Image)
	}
	assert.NotZero(t, len(d.Manga))
	for _, a := range d.Manga {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Image)
	}
	assert.NotZero(t, len(d.Character))
	for _, a := range d.Character {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Image)
	}
	assert.NotZero(t, len(d.People))
	for _, a := range d.People {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Name)
		assert.NotEmpty(t, a.Image)
	}
	time.Sleep(sleepDur)
}

func TestGetUserFriend(t *testing.T) {
	d, code, err := mal.GetUserFriend("rl404", 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg := true
	for _, f := range d {
		assert.NotEmpty(t, f.Username)
		if f.Image != "" {
			emptyImg = false
		}
		assert.NotZero(t, f.LastOnline)
		assert.NotZero(t, f.FriendSince)
	}
	assert.False(t, emptyImg)
	time.Sleep(sleepDur)
}

func TestGetUserHistory(t *testing.T) {
	d, code, err := mal.GetUserHistory("rl404", 0)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, h := range d {
		assert.NotZero(t, h.ID)
		assert.NotEmpty(t, h.Title)
		assert.NotEmpty(t, h.Type)
		assert.NotZero(t, h.Progress)
		assert.NotZero(t, h.Date)
	}
	time.Sleep(sleepDur)
}

func TestGetUserReview(t *testing.T) {
	d, code, err := mal.GetUserReview("Archaeon", 3)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyEp, emptyCh := true, true
	story0, anim0, sou0, char0, enj0 := true, true, true, true, true
	for _, r := range d {
		assert.NotZero(t, r.ID)
		assert.NotEmpty(t, r.Username)
		assert.NotEmpty(t, r.Image)
		assert.NotZero(t, r.Source.ID)
		assert.NotEmpty(t, r.Source.Title)
		assert.NotEmpty(t, r.Source.Image)
		assert.NotEmpty(t, r.Source.Type)
		assert.NotZero(t, r.Helpful)
		if r.Episode != "" {
			emptyEp = false
		}
		if r.Chapter != "" {
			emptyCh = false
		}
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
	assert.False(t, emptyEp)
	assert.False(t, emptyCh)
	assert.False(t, story0)
	assert.False(t, anim0)
	assert.False(t, sou0)
	assert.False(t, char0)
	assert.False(t, enj0)
	time.Sleep(sleepDur)
}

func TestGetUserRecommendation(t *testing.T) {
	d, code, err := mal.GetUserRecommendation("Archaeon", 1)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, r := range d {
		assert.NotZero(t, r.Source.ID)
		assert.NotEmpty(t, r.Source.Title)
		assert.NotEmpty(t, r.Source.Image)
		assert.NotEmpty(t, r.Source.Type)
		assert.NotZero(t, r.Recommended.ID)
		assert.NotEmpty(t, r.Recommended.Title)
		assert.NotEmpty(t, r.Recommended.Image)
		assert.NotEmpty(t, r.Recommended.Type)
		assert.Equal(t, len(r.Users), 1)
		for _, u := range r.Users {
			assert.NotEmpty(t, u.Username)
			assert.NotEmpty(t, u.Content)
		}
	}
	time.Sleep(sleepDur)
}

func TestGetUserClub(t *testing.T) {
	d, code, err := mal.GetUserClub("Archaeon")
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Name)
	}
	time.Sleep(sleepDur)
}

func TestGetUserAnime(t *testing.T) {
	d, code, err := mal.GetUserAnime("rl404", 0)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyScore, emptyProg, emptyEp, emptyTag, emptyRating := true, true, true, true, true, true
	for _, a := range d {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Title)
		if a.Image != "" {
			emptyImg = false
		}
		assert.NotZero(t, a.Status)
		if a.Score > 0 {
			emptyScore = false
		}
		assert.NotZero(t, a.Type)
		if a.Progress > 0 {
			emptyProg = false
		}
		if a.Episode > 0 {
			emptyEp = false
		}
		if a.Tag != "" {
			emptyTag = false
		}
		if a.Rating != "" {
			emptyRating = false
		}
		assert.NotZero(t, a.AiringStatus)
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyScore)
	assert.False(t, emptyProg)
	assert.False(t, emptyEp)
	assert.False(t, emptyTag)
	assert.False(t, emptyRating)
	time.Sleep(sleepDur)
}

func TestGetUserManga(t *testing.T) {
	d, code, err := mal.GetUserManga("rl404", 0)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyScore, emptyChProg, emptyVolProg, emptyCh, emptyVol, emptyTag := true, true, true, true, true, true, true
	for _, a := range d {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Title)
		if a.Image != "" {
			emptyImg = false
		}
		assert.NotZero(t, a.Status)
		if a.Score > 0 {
			emptyScore = false
		}
		assert.NotZero(t, a.Type)
		if a.ChapterProgress > 0 {
			emptyChProg = false
		}
		if a.VolumeProgress > 0 {
			emptyVolProg = false
		}
		if a.Chapter > 0 {
			emptyCh = false
		}
		if a.Volume > 0 {
			emptyVol = false
		}
		if a.Tag != "" {
			emptyTag = false
		}
		assert.NotZero(t, a.PublishingStatus)
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyScore)
	assert.False(t, emptyChProg)
	assert.False(t, emptyVolProg)
	assert.False(t, emptyCh)
	assert.False(t, emptyVol)
	assert.False(t, emptyTag)
	time.Sleep(sleepDur)
}
