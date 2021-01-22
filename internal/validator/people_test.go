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

func TestGetPeople(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetPeople(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeople(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetPeople", 1).Return(&model.People{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:people:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeople(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetPeopleCharacter(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetPeopleCharacter(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleCharacter(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetPeopleCharacter", 1).Return([]model.PeopleCharacter{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:people:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleCharacter(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetPeopleStaff(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetPeopleStaff(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleStaff(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetPeopleStaff", 1).Return([]model.Role{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:people:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleStaff(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetPeopleManga(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetPeopleManga(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleManga(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetPeopleManga", 1).Return([]model.Role{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:people:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleManga(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetPeopleNews(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetPeopleNews(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleNews(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetPeopleNews", 1).Return([]model.NewsItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:people:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleNews(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetPeopleArticle(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetPeopleArticle(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleArticle(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetPeopleArticle", 1).Return([]model.ArticleItem{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:people:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeopleArticle(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestGetPeoplePicture(t *testing.T) {
	var empty bool
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-id", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.GetPeoplePicture(0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("empty-id", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeoplePicture(1)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusNotFound, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Get", "mal:empty:people:1", &empty).Return(errDummy).Once()
		mockAPI.On("GetPeoplePicture", 1).Return([]string{}, http.StatusOK, nil).Once()
		mockCacher.On("Set", "mal:empty:people:1", true).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.GetPeoplePicture(1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}