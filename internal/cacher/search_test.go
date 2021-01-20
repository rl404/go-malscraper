package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearchAnime(t *testing.T) {
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockParser.On("SearchAnime", model.Query{}).Return([]model.AnimeSearch{}, http.StatusOK, nil).Once()
	c := Cacher{api: mockParser, cacher: mockCacher}

	d, code, err := c.SearchAnime(model.Query{})
	assert.NotNil(t, d)
	assert.Equal(t, code, http.StatusOK)
	assert.NoError(t, err)
}

func TestSearchManga(t *testing.T) {
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockParser.On("SearchManga", model.Query{}).Return([]model.MangaSearch{}, http.StatusOK, nil).Once()
	c := Cacher{api: mockParser, cacher: mockCacher}

	d, code, err := c.SearchManga(model.Query{})
	assert.NotNil(t, d)
	assert.Equal(t, code, http.StatusOK)
	assert.NoError(t, err)
}

func TestSearchCharacter(t *testing.T) {
	var data []model.CharacterSearch
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:search-character:name:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.CharacterSearch)
			*tmp = []model.CharacterSearch{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchCharacter("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("SearchCharacter", "name", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:search-character:name:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchCharacter("name", 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("SearchCharacter", "name", 1).Return([]model.CharacterSearch{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:search-character:name:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:search-character:name:1", []model.CharacterSearch{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchCharacter("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestSearchPeople(t *testing.T) {
	var data []model.PeopleSearch
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:search-people:name:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.PeopleSearch)
			*tmp = []model.PeopleSearch{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchPeople("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("SearchPeople", "name", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:search-people:name:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchPeople("name", 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("SearchPeople", "name", 1).Return([]model.PeopleSearch{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:search-people:name:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:search-people:name:1", []model.PeopleSearch{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchPeople("name", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestSearchClub(t *testing.T) {
	var data []model.ClubSearch
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:search-club:name:1:2:3", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ClubSearch)
			*tmp = []model.ClubSearch{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchClub(model.ClubQuery{Name: "name", Page: 1, Category: 2, Sort: 3})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("SearchClub", model.ClubQuery{Name: "name", Page: 1, Category: 2, Sort: 3}).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:search-club:name:1:2:3", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchClub(model.ClubQuery{Name: "name", Page: 1, Category: 2, Sort: 3})
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("SearchClub", model.ClubQuery{Name: "name", Page: 1, Category: 2, Sort: 3}).Return([]model.ClubSearch{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:search-club:name:1:2:3", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:search-club:name:1:2:3", []model.ClubSearch{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchClub(model.ClubQuery{Name: "name", Page: 1, Category: 2, Sort: 3})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestSearchUser(t *testing.T) {
	var data []model.UserSearch
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:search-user:name:1:loc:2:3:4", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.UserSearch)
			*tmp = []model.UserSearch{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchUser(model.UserQuery{Username: "name", Page: 1, Location: "loc", MinAge: 2, MaxAge: 3, Gender: 4})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("SearchUser", model.UserQuery{Username: "name", Page: 1, Location: "loc", MinAge: 2, MaxAge: 3, Gender: 4}).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:search-user:name:1:loc:2:3:4", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchUser(model.UserQuery{Username: "name", Page: 1, Location: "loc", MinAge: 2, MaxAge: 3, Gender: 4})
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("SearchUser", model.UserQuery{Username: "name", Page: 1, Location: "loc", MinAge: 2, MaxAge: 3, Gender: 4}).Return([]model.UserSearch{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:search-user:name:1:loc:2:3:4", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:search-user:name:1:loc:2:3:4", []model.UserSearch{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.SearchUser(model.UserQuery{Username: "name", Page: 1, Location: "loc", MinAge: 2, MaxAge: 3, Gender: 4})
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
