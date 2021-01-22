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

func TestGetRecommendation(t *testing.T) {
	var empty1, empty2 bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetRecommendation("", 0, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetRecommendation("anime", 0, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-anime-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty1).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		mockCacher.On("Get", "mal:empty:anime:2", &empty2).Return(errDummy).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetRecommendation("anime", 1, 2)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("empty-manga-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:manga:1", &empty1).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		mockCacher.On("Get", "mal:empty:manga:2", &empty2).Return(errDummy).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetRecommendation("manga", 1, 2)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty1).Return(errDummy).Once()
		mockCacher.On("Get", "mal:empty:anime:2", &empty2).Return(errDummy).Once()
		mockAPI.On("GetRecommendation", "anime", 1, 2).Return(&model.Recommendation{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetRecommendation("anime", 1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetRecommendations(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetRecommendations("", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetRecommendations("anime", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetRecommendations", "anime", 1).Return([]model.Recommendation{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetRecommendations("anime", 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}
