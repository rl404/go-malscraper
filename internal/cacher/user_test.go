package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	var data *model.User
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user:name", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.User)
			*tmp = &model.User{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUser("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUser", "name").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user:name", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUser("name")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUser", "name").Return(&model.User{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user:name", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user:name", &model.User{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUser("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserStats(t *testing.T) {
	var data *model.UserStats
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-stats:name", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.UserStats)
			*tmp = &model.UserStats{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserStats("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserStats", "name").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-stats:name", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserStats("name")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserStats", "name").Return(&model.UserStats{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-stats:name", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-stats:name", &model.UserStats{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserStats("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserFavorite(t *testing.T) {
	var data *model.UserFavorite
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-favorite:name", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.UserFavorite)
			*tmp = &model.UserFavorite{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserFavorite("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserFavorite", "name").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-favorite:name", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserFavorite("name")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserFavorite", "name").Return(&model.UserFavorite{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-favorite:name", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-favorite:name", &model.UserFavorite{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserFavorite("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserFriend(t *testing.T) {
	var data []model.UserFriend
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-friend:name:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.UserFriend)
			*tmp = []model.UserFriend{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserFriend("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserFriend", "name", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-friend:name:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserFriend("name", 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserFriend", "name", 1).Return([]model.UserFriend{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-friend:name:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-friend:name:1", []model.UserFriend{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserFriend("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserHistory(t *testing.T) {
	var data []model.UserHistory
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-history:name:type", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.UserHistory)
			*tmp = []model.UserHistory{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserHistory("name", "type")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserHistory", "name", "type").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-history:name:type", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserHistory("name", "type")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserHistory", "name", "type").Return([]model.UserHistory{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-history:name:type", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-history:name:type", []model.UserHistory{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserHistory("name", "type")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserReview(t *testing.T) {
	var data []model.Review
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-review:name:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Review)
			*tmp = []model.Review{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserReview("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserReview", "name", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-review:name:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserReview("name", 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserReview", "name", 1).Return([]model.Review{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-review:name:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-review:name:1", []model.Review{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserReview("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserRecommendation(t *testing.T) {
	var data []model.Recommendation
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-recommendation:name:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Recommendation)
			*tmp = []model.Recommendation{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserRecommendation("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserRecommendation", "name", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-recommendation:name:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserRecommendation("name", 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserRecommendation", "name", 1).Return([]model.Recommendation{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-recommendation:name:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-recommendation:name:1", []model.Recommendation{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserRecommendation("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserClub(t *testing.T) {
	var data []model.Item
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-club:name", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Item)
			*tmp = []model.Item{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserClub("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserClub", "name").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-club:name", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserClub("name")
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserClub", "name").Return([]model.Item{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-club:name", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-club:name", []model.Item{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserClub("name")
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserAnime(t *testing.T) {
	var data []model.UserAnime
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-anime:name:1:2:3:tag", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.UserAnime)
			*tmp = []model.UserAnime{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserAnime(model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserAnime", model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"}).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-anime:name:1:2:3:tag", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserAnime(model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"})
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserAnime", model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"}).Return([]model.UserAnime{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-anime:name:1:2:3:tag", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-anime:name:1:2:3:tag", []model.UserAnime{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserAnime(model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetUserManga(t *testing.T) {
	var data []model.UserManga
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:user-manga:name:1:2:3:tag", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.UserManga)
			*tmp = []model.UserManga{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserManga(model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetUserManga", model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"}).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:user-manga:name:1:2:3:tag", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserManga(model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"})
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetUserManga", model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"}).Return([]model.UserManga{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:user-manga:name:1:2:3:tag", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:user-manga:name:1:2:3:tag", []model.UserManga{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetUserManga(model.UserListQuery{Username: "name", Page: 1, Status: 2, Order: 3, Tag: "tag"})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
