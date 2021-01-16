package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetNews(t *testing.T) {
	var data *model.News
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:news:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.News)
			*tmp = &model.News{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetNews", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:news:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNews(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetNews", 1).Return(&model.News{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:news:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:news:1", &model.News{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetNewsList(t *testing.T) {
	var data []model.NewsItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:news-list:1:tag", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.NewsItem)
			*tmp = []model.NewsItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNewsList(1, "tag")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetNewsList", 1, "tag").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:news-list:1:tag", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNewsList(1, "tag")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetNewsList", 1, "tag").Return([]model.NewsItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:news-list:1:tag", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:news-list:1:tag", []model.NewsItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNewsList(1, "tag")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetNewsTag(t *testing.T) {
	var data *model.NewsTag
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:news-tag", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.NewsTag)
			*tmp = &model.NewsTag{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNewsTag()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetNewsTag").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:news-tag", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNewsTag()
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetNewsTag").Return(&model.NewsTag{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:news-tag", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:news-tag", &model.NewsTag{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetNewsTag()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
