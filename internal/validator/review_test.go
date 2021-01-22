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

func TestGetReview(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetReview(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:review:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetReview(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:review:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetReview", 1).Return(&model.Review{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:review:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetReview(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetReviews(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetReviews("", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetReviews("anime", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetReviews", "anime", 1).Return([]model.Review{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetReviews("anime", 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}
