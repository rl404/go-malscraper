package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTopAnime(t *testing.T) {
	var data []model.TopAnime
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:top-anime:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.TopAnime)
			*tmp = []model.TopAnime{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopAnime(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetTopAnime", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:top-anime:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopAnime(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetTopAnime", 1, 2).Return([]model.TopAnime{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:top-anime:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:top-anime:1:2", []model.TopAnime{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopAnime(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetTopManga(t *testing.T) {
	var data []model.TopManga
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:top-manga:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.TopManga)
			*tmp = []model.TopManga{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopManga(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetTopManga", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:top-manga:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopManga(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetTopManga", 1, 2).Return([]model.TopManga{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:top-manga:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:top-manga:1:2", []model.TopManga{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopManga(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetTopCharacter(t *testing.T) {
	var data []model.TopCharacter
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:top-character:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.TopCharacter)
			*tmp = []model.TopCharacter{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetTopCharacter", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:top-character:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopCharacter(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetTopCharacter", 1).Return([]model.TopCharacter{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:top-character:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:top-character:1", []model.TopCharacter{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetTopPeople(t *testing.T) {
	var data []model.TopPeople
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:top-people:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.TopPeople)
			*tmp = []model.TopPeople{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopPeople(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetTopPeople", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:top-people:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopPeople(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetTopPeople", 1).Return([]model.TopPeople{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:top-people:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:top-people:1", []model.TopPeople{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetTopPeople(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
