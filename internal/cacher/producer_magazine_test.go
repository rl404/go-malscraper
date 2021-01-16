package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProducers(t *testing.T) {
	var data []model.ItemCount
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:producers", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetProducers()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetProducers").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:producers", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetProducers()
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetProducers").Return([]model.ItemCount{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:producers", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:producers", []model.ItemCount{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetProducers()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetProducer(t *testing.T) {
	var data []model.AnimeItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:producer:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.AnimeItem)
			*tmp = []model.AnimeItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetProducer(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetProducer", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:producer:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetProducer(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetProducer", 1, 2).Return([]model.AnimeItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:producer:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:producer:1:2", []model.AnimeItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetProducer(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMagazines(t *testing.T) {
	var data []model.ItemCount
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:magazines", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMagazines()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMagazines").Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:magazines", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMagazines()
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMagazines").Return([]model.ItemCount{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:magazines", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:magazines", []model.ItemCount{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMagazines()
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetMagazine(t *testing.T) {
	var data []model.MangaItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:magazine:1:2", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.MangaItem)
			*tmp = []model.MangaItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMagazine(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetMagazine", 1, 2).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:magazine:1:2", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMagazine(1, 2)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetMagazine", 1, 2).Return([]model.MangaItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:magazine:1:2", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:magazine:1:2", []model.MangaItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetMagazine(1, 2)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
