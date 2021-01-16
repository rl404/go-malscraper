package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetGenres(t *testing.T) {
	var data []model.ItemCount
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:genres:type", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetGenres("type")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetGenres", "type").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:genres:type", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetGenres("type")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetGenres", "type").Return([]model.ItemCount{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:genres:type", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:genres:type", []model.ItemCount{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetGenres("type")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeWithGenre(t *testing.T) {
	var data []model.AnimeItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-with-genre:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.AnimeItem)
			*tmp = []model.AnimeItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeWithGenre(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeWithGenre", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-with-genre:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeWithGenre(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeWithGenre", 1, 2).Return([]model.AnimeItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-with-genre:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-with-genre:1:2", []model.AnimeItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeWithGenre(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaWithGenre(t *testing.T) {
	var data []model.MangaItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-with-genre:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.MangaItem)
			*tmp = []model.MangaItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaWithGenre(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaWithGenre", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-with-genre:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaWithGenre(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaWithGenre", 1, 2).Return([]model.MangaItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-with-genre:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-with-genre:1:2", []model.MangaItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaWithGenre(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
