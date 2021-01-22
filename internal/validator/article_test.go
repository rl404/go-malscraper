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

func TestGetArticle(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetArticle(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:article:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetArticle(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:article:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetArticle", 1).Return(&model.Article{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:article:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetArticles(t *testing.T) {
	var tags []model.ArticleTagItem
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetArticles(0, "")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-tag", func(t *testing.T) {
		mockCacher.On("Get", "mal:article-tag", &tags).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleTagItem)
			*tmp = []model.ArticleTagItem{}
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetArticles(1, "tag")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidTag.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("GetArticles", 1, "").Return([]model.ArticleItem{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetArticles(1, "")
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetArticleTag(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	mockAPI.On("GetArticleTag").Return([]model.ArticleTagItem{}, http.StatusOK, nil).Once()
	v := New(mockAPI, mockCacher, mockLogger)

	d, code, err := v.GetArticleTag()
	assert.NotNil(t, d)
	assert.Equal(t, http.StatusOK, code)
	assert.NoError(t, err)
}
