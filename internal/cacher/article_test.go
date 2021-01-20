package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetArticle(t *testing.T) {
	var data *model.Article
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:article:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Article)
			*tmp = &model.Article{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetArticle", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:article:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticle(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetArticle", 1).Return(&model.Article{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:article:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:article:1", &model.Article{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetArticles(t *testing.T) {
	var data []model.ArticleItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:article-list:1:tag", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleItem)
			*tmp = []model.ArticleItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticles(1, "tag")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetArticles", 1, "tag").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:article-list:1:tag", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticles(1, "tag")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetArticles", 1, "tag").Return([]model.ArticleItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:article-list:1:tag", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:article-list:1:tag", []model.ArticleItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticles(1, "tag")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetArticleTag(t *testing.T) {
	var data []model.ArticleTagItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:article-tag", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleTagItem)
			*tmp = []model.ArticleTagItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticleTag()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetArticleTag").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:article-tag", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticleTag()
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetArticleTag").Return([]model.ArticleTagItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:article-tag", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:article-tag", []model.ArticleTagItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetArticleTag()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
