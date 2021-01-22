package validator

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service/mocks"
	"github.com/rl404/mal-plugin/log/mallogger"
	"github.com/stretchr/testify/assert"
)

func TestGetTopAnime(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetTopAnime(-1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetTopAnime(1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetTopAnime", 1, 2).Return([]model.TopAnime{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetTopAnime(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetTopManga(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetTopManga(-1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetTopManga(1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetTopManga", 1, 2).Return([]model.TopManga{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetTopManga(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetTopCharacter(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetTopCharacter(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetTopCharacter", 1).Return([]model.TopCharacter{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetTopCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetTopPeople(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetTopPeople(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetTopPeople", 1).Return([]model.TopPeople{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetTopPeople(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}
