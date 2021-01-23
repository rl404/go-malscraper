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

func TestSearchAnime(t *testing.T) {
	var producers, genres []model.ItemCount
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-title", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchAnime(model.Query{Title: ""})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.Err3LettersSearch.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchAnime(model.Query{Title: "naruto", Page: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchAnime(model.Query{Title: "naruto", Type: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("invalid-score", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchAnime(model.Query{Title: "naruto", Score: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidScore.Error())
	})

	t.Run("invalid-status", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchAnime(model.Query{Title: "naruto", Status: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidStatus.Error())
	})

	t.Run("invalid-producer", func(t *testing.T) {
		mockCacher.On("Get", "mal:producers", &producers).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchAnime(model.Query{Title: "naruto", ProducerID: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidProducer.Error())
	})

	t.Run("invalid-genre", func(t *testing.T) {
		mockCacher.On("Get", "mal:genres:anime", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchAnime(model.Query{Title: "naruto", GenreIDs: []int{0}})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidGenre.Error())
	})

	t.Run("invalid-rating", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchAnime(model.Query{Title: "naruto", Rating: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidRating.Error())
	})

	t.Run("invalid-firstletter", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchAnime(model.Query{Title: "naruto", FirstLetter: "ab"})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidFirstLetter.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("SearchAnime", model.Query{Title: "naruto", Page: 1}).Return([]model.AnimeSearch{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchAnime(model.Query{Title: "naruto"})
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestSearchManga(t *testing.T) {
	var magazines, genres []model.ItemCount
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-title", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchManga(model.Query{Title: ""})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.Err3LettersSearch.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchManga(model.Query{Title: "naruto", Page: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-type", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchManga(model.Query{Title: "naruto", Type: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidType.Error())
	})

	t.Run("invalid-score", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchManga(model.Query{Title: "naruto", Score: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidScore.Error())
	})

	t.Run("invalid-status", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchManga(model.Query{Title: "naruto", Status: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidStatus.Error())
	})

	t.Run("invalid-magazine", func(t *testing.T) {
		mockCacher.On("Get", "mal:magazines", &magazines).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchManga(model.Query{Title: "naruto", MagazineID: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidMagazine.Error())
	})

	t.Run("invalid-genre", func(t *testing.T) {
		mockCacher.On("Get", "mal:genres:manga", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{}
		}).Return(nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchManga(model.Query{Title: "naruto", GenreIDs: []int{0}})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidGenre.Error())
	})

	t.Run("invalid-firstletter", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchManga(model.Query{Title: "naruto", FirstLetter: "ab"})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidFirstLetter.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("SearchManga", model.Query{Title: "naruto", Page: 1}).Return([]model.MangaSearch{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchManga(model.Query{Title: "naruto"})
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestSearchCharacter(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-name", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchCharacter("", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.Err3LettersSearch.Error())
	})

	t.Run("invalid-name", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchCharacter("naruto", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("SearchCharacter", "naruto", 1).Return([]model.CharacterSearch{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchCharacter("naruto", 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestSearchPeople(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-name", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchPeople("", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.Err3LettersSearch.Error())
	})

	t.Run("invalid-name", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchPeople("naruto", 0)
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("SearchPeople", "naruto", 1).Return([]model.PeopleSearch{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchPeople("naruto", 1)
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestSearchClub(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-name", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchClub(model.ClubQuery{Name: ""})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.Err3LettersSearch.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchClub(model.ClubQuery{Name: "naruto", Page: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-category", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchClub(model.ClubQuery{Name: "naruto", Category: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidClubCategory.Error())
	})

	t.Run("invalid-sort", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchClub(model.ClubQuery{Name: "naruto", Sort: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidSortType.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("SearchClub", model.ClubQuery{Name: "naruto", Page: 1}).Return([]model.ClubSearch{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchClub(model.ClubQuery{Name: "naruto"})
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}

func TestSearchUser(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := mallogger.New(0, false)

	t.Run("invalid-name", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchUser(model.UserQuery{Username: ""})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.Err3LettersSearch.Error())
	})

	t.Run("invalid-page", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchUser(model.UserQuery{Username: "rl404", Page: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidPage.Error())
	})

	t.Run("invalid-age", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchUser(model.UserQuery{Username: "rl404", MinAge: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidAge.Error())
	})

	t.Run("invalid-gender", func(t *testing.T) {
		v := New(mockAPI, mockCacher, mockLogger)
		d, code, err := v.SearchUser(model.UserQuery{Username: "rl404", Gender: -1})
		assert.Nil(t, d)
		assert.Equal(t, http.StatusBadRequest, code)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInvalidGender.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockAPI.On("SearchUser", model.UserQuery{Username: "rl404", Page: 1}).Return([]model.UserSearch{}, http.StatusOK, nil).Once()
		v := New(mockAPI, mockCacher, mockLogger)

		d, code, err := v.SearchUser(model.UserQuery{Username: "rl404"})
		assert.NotNil(t, d)
		assert.Equal(t, http.StatusOK, code)
		assert.NoError(t, err)
	})
}
