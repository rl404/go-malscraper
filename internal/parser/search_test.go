package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchAnime(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.SearchAnime(model.Query{
		Title:        "naruto",
		Page:         1,
		Type:         1,
		Score:        7,
		Status:       2,
		ExcludeGenre: true,
		GenreIDs:     []int{1},
	})

	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	for _, a := range d {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, a.Image)
		assert.NotEmpty(t, a.Summary)
		assert.NotEmpty(t, a.Type)
		assert.NotZero(t, a.Episode)
		assert.NotZero(t, a.Score)
		assert.NotZero(t, a.StartDate.Year)
		assert.NotZero(t, a.StartDate.Month)
		assert.NotZero(t, a.StartDate.Day)
		assert.NotZero(t, a.EndDate.Year)
		assert.NotZero(t, a.EndDate.Month)
		assert.NotZero(t, a.EndDate.Day)
		assert.NotZero(t, a.Member)
		assert.NotEmpty(t, a.Rated)
	}
	time.Sleep(sleepDur)
}

func TestSearchManga(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.SearchManga(model.Query{
		Title:        "naruto",
		Page:         1,
		Type:         1,
		Score:        7,
		Status:       2,
		ExcludeGenre: true,
		GenreIDs:     []int{1},
	})

	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptySum, emptyVol, emptyCh := true, true, true
	emptyY1, emptyM1, emptyD1, emptyY2, emptyM2, emptyD2 := true, true, true, true, true, true
	for _, a := range d {
		assert.NotZero(t, a.ID)
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, a.Image)
		if a.Summary != "" {
			emptySum = false
		}
		assert.NotEmpty(t, a.Type)
		if a.Volume > 0 {
			emptyVol = false
		}
		if a.Chapter > 0 {
			emptyCh = false
		}
		assert.NotZero(t, a.Score)
		if a.StartDate.Year != 0 {
			emptyY1 = false
		}
		if a.StartDate.Month != 0 {
			emptyM1 = false
		}
		if a.StartDate.Day != 0 {
			emptyD1 = false
		}
		if a.EndDate.Year != 0 {
			emptyY2 = false
		}
		if a.EndDate.Month != 0 {
			emptyM2 = false
		}
		if a.EndDate.Day != 0 {
			emptyD2 = false
		}
		assert.NotZero(t, a.Member)
	}
	assert.False(t, emptySum)
	assert.False(t, emptyVol)
	assert.False(t, emptyCh)
	assert.False(t, emptyY1)
	assert.False(t, emptyM1)
	assert.False(t, emptyD1)
	assert.False(t, emptyY2)
	assert.False(t, emptyM2)
	assert.False(t, emptyD2)
	time.Sleep(sleepDur)
}

func TestSearchCharacter(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.SearchCharacter("naruto", 1)

	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyNick := true, true
	for _, c := range d {
		assert.NotZero(t, c.ID)
		assert.NotEmpty(t, c.Name)
		if c.Image != "" {
			emptyImg = false
		}
		if c.Nickname != "" {
			emptyNick = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyNick)
	time.Sleep(sleepDur)
}

func TestSearchPeople(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.SearchPeople("kana", 1)

	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyNick := true, true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Nickname != "" {
			emptyNick = false
		}
		if p.Image != "" {
			emptyImg = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyNick)
	time.Sleep(sleepDur)
}

func TestSearchClub(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.SearchClub(model.ClubQuery{
		Name:     "one",
		Page:     1,
		Category: 1,
		Sort:     1,
	})

	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptyMem := true, true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Image != "" {
			emptyImg = false
		}
		assert.NotEmpty(t, p.Summary)
		assert.NotEmpty(t, p.Creator)
		if p.Member > 0 {
			emptyMem = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptyMem)
	time.Sleep(sleepDur)
}

func TestSearchUser(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.SearchUser(model.UserQuery{
		Username: "404",
		Page:     1,
		Location: "California",
		MinAge:   20,
		MaxAge:   30,
		Gender:   1,
	})

	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg := true
	for _, p := range d {
		assert.NotEmpty(t, p.Username)
		if p.Image != "" {
			emptyImg = false
		}
		assert.NotNil(t, p.LastOnline)
	}
	assert.False(t, emptyImg)
}
