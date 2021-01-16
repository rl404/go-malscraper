package cacher

import (
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCharacter(t *testing.T) {
	var data *model.Character
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:character:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(**model.Character)
			*tmp = &model.Character{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetCharacter", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:character:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacter(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetCharacter", 1).Return(&model.Character{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:character:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:character:1", &model.Character{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetCharacterArticle(t *testing.T) {
	var data []model.ArticleItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:character-article:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleItem)
			*tmp = []model.ArticleItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetCharacterArticle", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:character-article:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterArticle(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetCharacterArticle", 1).Return([]model.ArticleItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:character-article:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:character-article:1", []model.ArticleItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetCharacterOgraphy(t *testing.T) {
	var data []model.Role
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:character-ography:type:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.Role)
			*tmp = []model.Role{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterOgraphy("type", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetCharacterOgraphy", "type", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:character-ography:type:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterOgraphy("type", 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetCharacterOgraphy", "type", 1).Return([]model.Role{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:character-ography:type:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:character-ography:type:1", []model.Role{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterOgraphy("type", 1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetCharacterPicture(t *testing.T) {
	var data []string
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:character-picture:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]string)
			*tmp = []string{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterPicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetCharacterPicture", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:character-picture:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterPicture(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetCharacterPicture", 1).Return([]string{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:character-picture:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:character-picture:1", []string{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterPicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetCharacterClub(t *testing.T) {
	var data []model.ClubItem
	mockParser := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	t.Run("cached", func(t *testing.T) {
		mockCacher.On("Get", "mal:character-club:1", &data).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ClubItem)
			*tmp = []model.ClubItem{}
		}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterClub(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockParser.On("GetCharacterClub", 1).Return(nil, http.StatusInternalServerError, errDummy).Once()
		mockCacher.On("Get", "mal:character-club:1", &data).Return(errDummy).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterClub(1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockParser.On("GetCharacterClub", 1).Return([]model.ClubItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Get", "mal:character-club:1", &data).Return(errDummy).Once()
		mockCacher.On("Set", "mal:character-club:1", []model.ClubItem{}).Return(nil).Once()
		c := Cacher{api: mockParser, cacher: mockCacher}

		d, code, err := c.GetCharacterClub(1)
		assert.NotNil(t, d)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}
