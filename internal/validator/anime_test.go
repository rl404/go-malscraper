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

func TestGetAnime(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnime(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnime(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnime", 1).Return(&model.Anime{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnime(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeCharacter(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeCharacter(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeCharacter(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeCharacter", 1).Return([]model.CharacterItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeStaff(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeStaff(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeStaff(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeStaff", 1).Return([]model.Role{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeStaff(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeVideo(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeVideo(0, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeVideo(1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeVideo(1, 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeVideo", 1, 1).Return(&model.Video{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeVideo(1, 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeEpisode(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeEpisode(0, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeEpisode(1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeEpisode(1, 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeEpisode", 1, 1).Return([]model.Episode{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeEpisode(1, 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeStats(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeStats(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeStats(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeStats", 1).Return(&model.Stats{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeStats(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeReview(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeReview(0, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeReview(1, 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeReview(1, 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeReview", 1, 1).Return([]model.Review{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeReview(1, 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeRecommendation(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeRecommendation(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeRecommendation(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeRecommendation", 1).Return([]model.Recommendation{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeRecommendation(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeNews(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeNews(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeNews(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeNews", 1).Return([]model.NewsItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeArticles(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeArticle(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeArticle(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeArticle", 1).Return([]model.ArticleItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeClub(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeClub(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeClub(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeClub", 1).Return([]model.ClubItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeClub(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimePicture(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimePicture(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimePicture(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimePicture", 1).Return([]string{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimePicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetAnimeMoreInfo(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetAnimeMoreInfo(0)
		assert.Empty(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeMoreInfo(1)
		assert.Empty(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:anime:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetAnimeMoreInfo", 1).Return("info", http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:anime:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetAnimeMoreInfo(1)
		assert.NotEmpty(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}