package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAnime(t *testing.T) {
	var data *model.Anime
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Anime)
			*tmp = &model.Anime{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnime(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnime", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnime(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnime", 1).Return(&model.Anime{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime:1", &model.Anime{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnime(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeCharacter(t *testing.T) {
	var data []model.CharacterItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-character:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.CharacterItem)
			*tmp = []model.CharacterItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeCharacter", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-character:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeCharacter(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeCharacter", 1).Return([]model.CharacterItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-character:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-character:1", []model.CharacterItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeStaff(t *testing.T) {
	var data []model.Role
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-staff:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Role)
			*tmp = []model.Role{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeStaff(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeStaff", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-staff:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeStaff(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeStaff", 1).Return([]model.Role{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-staff:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-staff:1", []model.Role{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeStaff(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeVideo(t *testing.T) {
	var data *model.Video
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-video:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Video)
			*tmp = &model.Video{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeVideo(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeVideo", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-video:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeVideo(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeVideo", 1, 2).Return(&model.Video{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-video:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-video:1:2", &model.Video{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeVideo(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeEpisode(t *testing.T) {
	var data []model.Episode
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-episode:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Episode)
			*tmp = []model.Episode{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeEpisode(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeEpisode", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-episode:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeEpisode(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeEpisode", 1, 2).Return([]model.Episode{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-episode:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-episode:1:2", []model.Episode{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeEpisode(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeStats(t *testing.T) {
	var data *model.Stats
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-stats:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Stats)
			*tmp = &model.Stats{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeStats(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeStats", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-stats:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeStats(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeStats", 1).Return(&model.Stats{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-stats:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-stats:1", &model.Stats{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeStats(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeReview(t *testing.T) {
	var data []model.Review
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-review:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Review)
			*tmp = []model.Review{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeReview(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeReview", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-review:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeReview(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeReview", 1, 2).Return([]model.Review{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-review:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-review:1:2", []model.Review{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeReview(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeRecommendation(t *testing.T) {
	var data []model.Recommendation
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-recommendation:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Recommendation)
			*tmp = []model.Recommendation{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeRecommendation(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeRecommendation", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-recommendation:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeRecommendation(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeRecommendation", 1).Return([]model.Recommendation{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-recommendation:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-recommendation:1", []model.Recommendation{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeRecommendation(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeNews(t *testing.T) {
	var data []model.NewsItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-news:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.NewsItem)
			*tmp = []model.NewsItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeNews", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-news:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeNews(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeNews", 1).Return([]model.NewsItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-news:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-news:1", []model.NewsItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeArticle(t *testing.T) {
	var data []model.ArticleItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-article:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleItem)
			*tmp = []model.ArticleItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeArticle", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-article:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeArticle(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeArticle", 1).Return([]model.ArticleItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-article:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-article:1", []model.ArticleItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeClub(t *testing.T) {
	var data []model.ClubItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-club:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ClubItem)
			*tmp = []model.ClubItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeClub(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeClub", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-club:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeClub(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeClub", 1).Return([]model.ClubItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-club:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-club:1", []model.ClubItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeClub(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimePicture(t *testing.T) {
	var data []string
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-picture:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]string)
			*tmp = []string{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimePicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimePicture", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-picture:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimePicture(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimePicture", 1).Return([]string{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-picture:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-picture:1", []string{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimePicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetAnimeMoreInfo(t *testing.T) {
	var data string
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:anime-more-info:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*string)
			*tmp = ""
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeMoreInfo(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetAnimeMoreInfo", 1).Return("", http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:anime-more-info:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeMoreInfo(1)
		assert.Empty(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetAnimeMoreInfo", 1).Return("string", http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:anime-more-info:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:anime-more-info:1", "string").Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetAnimeMoreInfo(1)
		assert.NotEmpty(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
