package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRecommendation(t *testing.T) {
	var data *model.Recommendation
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:recommendation:type:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Recommendation)
			*tmp = &model.Recommendation{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetRecommendation("type", 1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetRecommendation", "type", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:recommendation:type:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetRecommendation("type", 1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetRecommendation", "type", 1, 2).Return(&model.Recommendation{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:recommendation:type:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:recommendation:type:1:2", &model.Recommendation{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetRecommendation("type", 1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetRecommendations(t *testing.T) {
	var data []model.Recommendation
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:recommendations:type:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Recommendation)
			*tmp = []model.Recommendation{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetRecommendations("type", 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetRecommendations", "type", 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:recommendations:type:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetRecommendations("type", 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetRecommendations", "type", 2).Return([]model.Recommendation{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:recommendations:type:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:recommendations:type:2", []model.Recommendation{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetRecommendations("type", 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
