package validator

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service/mocks"
	"github.com/rl404/mal-plugin/log/mallogger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetGenres(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetGenres("")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetGenres", AnimeType).Return([]model.ItemCount{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetGenres(AnimeType)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeWithGenre(t *testing.T) {
	var genres []model.ItemCount
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeWithGenre(1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:genres:anime", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeWithGenre(1, 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:genres:anime", &genres).Return(errDummy).Once()
		mockAPI.On("GetAnimeWithGenre", 1, 1).Return([]model.AnimeItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:genres:anime", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeWithGenre(1, 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetMangaWithGenre(t *testing.T) {
	var genres []model.ItemCount
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetMangaWithGenre(1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:genres:manga", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetMangaWithGenre(1, 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:genres:manga", &genres).Return(errDummy).Once()
		mockAPI.On("GetMangaWithGenre", 1, 1).Return([]model.MangaItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:genres:manga", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetMangaWithGenre(1, 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}
