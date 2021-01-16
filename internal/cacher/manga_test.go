package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetManga(t *testing.T) {
	var data *model.Manga
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Manga)
			*tmp = &model.Manga{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetManga(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetManga", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetManga(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetManga", 1).Return(&model.Manga{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga:1", &model.Manga{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetManga(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaReview(t *testing.T) {
	var data []model.Review
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-review:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Review)
			*tmp = []model.Review{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaReview(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaReview", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-review:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaReview(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaReview", 1, 2).Return([]model.Review{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-review:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-review:1:2", []model.Review{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaReview(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaRecommendation(t *testing.T) {
	var data []model.Recommendation
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-recommendation:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Recommendation)
			*tmp = []model.Recommendation{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaRecommendation(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaRecommendation", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-recommendation:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaRecommendation(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaRecommendation", 1).Return([]model.Recommendation{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-recommendation:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-recommendation:1", []model.Recommendation{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaRecommendation(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaStats(t *testing.T) {
	var data *model.Stats
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-stats:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Stats)
			*tmp = &model.Stats{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaStats(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaStats", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-stats:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaStats(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaStats", 1).Return(&model.Stats{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-stats:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-stats:1", &model.Stats{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaStats(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaCharacter(t *testing.T) {
	var data []model.Role
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-character:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Role)
			*tmp = []model.Role{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaCharacter", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-character:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaCharacter(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaCharacter", 1).Return([]model.Role{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-character:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-character:1", []model.Role{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaNews(t *testing.T) {
	var data []model.NewsItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-news:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.NewsItem)
			*tmp = []model.NewsItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaNews", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-news:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaNews(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaNews", 1).Return([]model.NewsItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-news:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-news:1", []model.NewsItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaArticle(t *testing.T) {
	var data []model.ArticleItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-article:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleItem)
			*tmp = []model.ArticleItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaArticle", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-article:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaArticle(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaArticle", 1).Return([]model.ArticleItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-article:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-article:1", []model.ArticleItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaClub(t *testing.T) {
	var data []model.ClubItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-club:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ClubItem)
			*tmp = []model.ClubItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaClub(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaClub", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-club:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaClub(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaClub", 1).Return([]model.ClubItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-club:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-club:1", []model.ClubItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaClub(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaPicture(t *testing.T) {
	var data []string
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-picture:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]string)
			*tmp = []string{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaPicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaPicture", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-picture:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaPicture(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaPicture", 1).Return([]string{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-picture:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-picture:1", []string{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaPicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMangaMoreInfo(t *testing.T) {
	var data string
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:manga-more-info:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*string)
			*tmp = ""
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaMoreInfo(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMangaMoreInfo", 1).Return("", http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:manga-more-info:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaMoreInfo(1)
		assert.Empty(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMangaMoreInfo", 1).Return("string", http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:manga-more-info:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:manga-more-info:1", "string").Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMangaMoreInfo(1)
		assert.NotEmpty(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
