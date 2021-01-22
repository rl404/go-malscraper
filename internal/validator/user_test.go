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

func TestGetUser(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUser("")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUser("rl404")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUser", "rl404").Return(&model.User{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUser("rl404")
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserStats(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserStats("")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserStats("rl404")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserStats", "rl404").Return(&model.UserStats{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserStats("rl404")
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserFavorite(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserFavorite("")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserFavorite("rl404")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserFavorite", "rl404").Return(&model.UserFavorite{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserFavorite("rl404")
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserFriend(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserFriend("", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserFriend("rl404", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserFriend("rl404", 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserFriend", "rl404", 1).Return([]model.UserFriend{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserFriend("rl404", 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserHistory(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserHistory("", "")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserHistory("rl404", "t")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserHistory("rl404", "")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserHistory", "rl404", "").Return([]model.UserHistory{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserHistory("rl404", "")
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserReview(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserReview("", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserReview("rl404", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserReview("rl404", 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserReview", "rl404", 1).Return([]model.Review{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserReview("rl404", 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserRecommendation(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserRecommendation("", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserRecommendation("rl404", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserRecommendation("rl404", 1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserRecommendation", "rl404", 1).Return([]model.Recommendation{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserRecommendation("rl404", 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserClub(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserClub("")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserClub("rl404")
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserClub", "rl404").Return([]model.Item{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserClub("rl404")
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserAnime(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserAnime(model.UserListQuery{Username: ""})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserAnime(model.UserListQuery{Username: "rl404", Page: -2})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-status", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserAnime(model.UserListQuery{Username: "rl404", Status: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidStatus.Error())
	})

	t.Run("invalid-order", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserAnime(model.UserListQuery{Username: "rl404", Order: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidOrder.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserAnime(model.UserListQuery{Username: "rl404"})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserAnime", model.UserListQuery{Username: "rl404", Page: 1}).Return([]model.UserAnime{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserAnime(model.UserListQuery{Username: "rl404"})
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetUserManga(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-user", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserManga(model.UserListQuery{Username: ""})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidUsername.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserManga(model.UserListQuery{Username: "rl404", Page: -2})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-status", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserManga(model.UserListQuery{Username: "rl404", Status: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidStatus.Error())
	})

	t.Run("invalid-order", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetUserManga(model.UserListQuery{Username: "rl404", Order: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidOrder.Error())
	})

	t.Run("empty-user", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserManga(model.UserListQuery{Username: "rl404"})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:user:rl404", &empty).Return(errDummy).Once()
		mockAPI.On("GetUserManga", model.UserListQuery{Username: "rl404", Page: 1}).Return([]model.UserManga{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:user:rl404", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetUserManga(model.UserListQuery{Username: "rl404"})
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}
