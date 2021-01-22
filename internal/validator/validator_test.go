package validator

import (
	"errors"
	"net/http"
	"testing"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var errDummy = errors.New("dummy error")

func TestNew(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)
	_ = New(mockAPI, mockCacher, mockLogger)
}

func TestIsEmptyID(t *testing.T) {
	var empty bool
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("empty", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking empty id...", "key").Once()
		mockCacher.On("Get", "key", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(nil).Once()
		mockLogger.On("Debug", "[%s] found empty id", "key").Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}

		e := v.isEmptyID("key")
		assert.True(t, e)
	})

	t.Run("not-empty", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking empty id...", "key").Once()
		mockCacher.On("Get", "key", &empty).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*bool)
			*tmp = true
		}).Return(errDummy).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}

		e := v.isEmptyID("key")
		assert.False(t, e)
	})

}

func TestSaveEmptyID(t *testing.T) {
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("ok", func(t *testing.T) {
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		v.saveEmptyID(http.StatusOK, "key")
	})

	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] saving empty id...", "key")
		mockCacher.On("Set", "key", true).Return(errDummy).Once()
		mockLogger.On("Error", "[%s] failed saving cache: %s", "key", errDummy.Error())
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		v.saveEmptyID(http.StatusNotFound, "key")
	})

	t.Run("no-error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] saving empty id...", "key")
		mockCacher.On("Set", "key", true).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		v.saveEmptyID(http.StatusNotFound, "key")
	})
}

func TestIsArticleTagValid(t *testing.T) {
	var tags []model.ArticleTagItem
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("empty", func(t *testing.T) {
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isArticleTagValid("")
		assert.True(t, b)
	})

	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid article tag...", "mal:article-tag").Once()
		mockCacher.On("Get", "mal:article-tag", &tags).Return(errDummy).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isArticleTagValid("tag")
		assert.True(t, b)
	})

	t.Run("valid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid article tag...", "mal:article-tag").Once()
		mockCacher.On("Get", "mal:article-tag", &tags).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleTagItem)
			*tmp = []model.ArticleTagItem{{Tag: "tag"}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isArticleTagValid("tag")
		assert.True(t, b)
	})

	t.Run("invalid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid article tag...", "mal:article-tag").Once()
		mockCacher.On("Get", "mal:article-tag", &tags).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ArticleTagItem)
			*tmp = []model.ArticleTagItem{{Tag: "tag"}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isArticleTagValid("tag2")
		assert.False(t, b)
	})
}

func TestIsAnimeGenreValid(t *testing.T) {
	var genres []model.ItemCount
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("empty", func(t *testing.T) {
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isAnimeGenreValid(0)
		assert.False(t, b)
	})

	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid anime genre...", "mal:genres:anime").Once()
		mockCacher.On("Get", "mal:genres:anime", &genres).Return(errDummy).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isAnimeGenreValid(1)
		assert.True(t, b)
	})

	t.Run("valid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid anime genre...", "mal:genres:anime").Once()
		mockCacher.On("Get", "mal:genres:anime", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isAnimeGenreValid(1)
		assert.True(t, b)
	})

	t.Run("invalid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid anime genre...", "mal:genres:anime").Once()
		mockCacher.On("Get", "mal:genres:anime", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isAnimeGenreValid(2)
		assert.False(t, b)
	})
}

func TestIsMangaGenreValid(t *testing.T) {
	var genres []model.ItemCount
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("empty", func(t *testing.T) {
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMangaGenreValid(0)
		assert.False(t, b)
	})

	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid manga genre...", "mal:genres:manga").Once()
		mockCacher.On("Get", "mal:genres:manga", &genres).Return(errDummy).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMangaGenreValid(1)
		assert.True(t, b)
	})

	t.Run("valid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid manga genre...", "mal:genres:manga").Once()
		mockCacher.On("Get", "mal:genres:manga", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMangaGenreValid(1)
		assert.True(t, b)
	})

	t.Run("invalid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid manga genre...", "mal:genres:manga").Once()
		mockCacher.On("Get", "mal:genres:manga", &genres).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMangaGenreValid(2)
		assert.False(t, b)
	})
}

func TestIsNewsTagValid(t *testing.T) {
	var tags model.NewsTag
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("empty", func(t *testing.T) {
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isNewsTagValid("")
		assert.True(t, b)
	})

	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid news tag...", "mal:news-tag").Once()
		mockCacher.On("Get", "mal:news-tag", &tags).Return(errDummy).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isNewsTagValid("tag")
		assert.True(t, b)
	})

	t.Run("valid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid news tag...", "mal:news-tag")
		mockCacher.On("Get", "mal:news-tag", &tags).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*model.NewsTag)
			*tmp = model.NewsTag{
				Anime:    []model.NewsTagItem{{Tag: "tag1"}},
				Manga:    []model.NewsTagItem{{Tag: "tag2"}},
				People:   []model.NewsTagItem{{Tag: "tag3"}},
				Music:    []model.NewsTagItem{{Tag: "tag4"}},
				Event:    []model.NewsTagItem{{Tag: "tag5"}},
				Industry: []model.NewsTagItem{{Tag: "tag6"}},
			}
		}).Return(nil)
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		assert.True(t, v.isNewsTagValid("tag1"))
		assert.True(t, v.isNewsTagValid("tag2"))
		assert.True(t, v.isNewsTagValid("tag3"))
		assert.True(t, v.isNewsTagValid("tag4"))
		assert.True(t, v.isNewsTagValid("tag5"))
		assert.True(t, v.isNewsTagValid("tag6"))
	})

	t.Run("invalid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid news tag...", "mal:news-tag").Once()
		mockCacher.On("Get", "mal:news-tag", &tags).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*model.NewsTag)
			*tmp = model.NewsTag{}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isNewsTagValid("tag")
		assert.False(t, b)
	})
}

func TestIsProducerValid(t *testing.T) {
	var producers []model.ItemCount
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("empty", func(t *testing.T) {
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isProducerValid(-1)
		assert.False(t, b)
	})

	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid producer...", "mal:producers").Once()
		mockCacher.On("Get", "mal:producers", &producers).Return(errDummy).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isProducerValid(1)
		assert.True(t, b)
	})

	t.Run("valid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid producer...", "mal:producers").Once()
		mockCacher.On("Get", "mal:producers", &producers).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isProducerValid(1)
		assert.True(t, b)
	})

	t.Run("invalid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid producer...", "mal:producers").Once()
		mockCacher.On("Get", "mal:producers", &producers).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isProducerValid(2)
		assert.False(t, b)
	})
}

func TestIsMagazineValid(t *testing.T) {
	var magazine []model.ItemCount
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)

	t.Run("empty", func(t *testing.T) {
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMagazineValid(-1)
		assert.False(t, b)
	})

	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid magazine...", "mal:magazines").Once()
		mockCacher.On("Get", "mal:magazines", &magazine).Return(errDummy).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMagazineValid(1)
		assert.True(t, b)
	})

	t.Run("valid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid magazine...", "mal:magazines").Once()
		mockCacher.On("Get", "mal:magazines", &magazine).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMagazineValid(1)
		assert.True(t, b)
	})

	t.Run("invalid", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] checking valid magazine...", "mal:magazines").Once()
		mockCacher.On("Get", "mal:magazines", &magazine).Run(func(args mock.Arguments) {
			tmp := args.Get(1).(*[]model.ItemCount)
			*tmp = []model.ItemCount{{ID: 1}}
		}).Return(nil).Once()
		v := &Validator{cacher: mockCacher, logger: mockLogger}
		b := v.isMagazineValid(2)
		assert.False(t, b)
	})
}