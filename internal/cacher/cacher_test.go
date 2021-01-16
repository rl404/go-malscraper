package cacher

import (
	"errors"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/internal/mocks"
	"github.com/stretchr/testify/assert"
)

var errDummy = errors.New("dummy error")

func init() {
	timeSince = func(time.Time) time.Duration {
		return time.Second
	}
}

func TestNew(t *testing.T) {
	mockAPI := new(mocks.API)
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)
	_ = New(mockAPI, mockCacher, mockLogger)
}

func TestGet(t *testing.T) {
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)
	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] retrieving cache...", "key").Once()
		mockCacher.On("Get", "key", "data").Return(errDummy).Once()
		mockLogger.On("Warn", "[%s] failed retrieving cache: %s", "key", errDummy.Error()).Once()
		c := newCacherLog(mockCacher, mockLogger)

		err := c.Get("key", "data")
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] retrieving cache...", "key").Once()
		mockCacher.On("Get", "key", "data").Return(nil).Once()
		mockLogger.On("Debug", "[%s] cache found (%s)", "key", time.Second.Truncate(time.Microsecond)).Once()
		c := newCacherLog(mockCacher, mockLogger)

		err := c.Get("key", "data")
		assert.NoError(t, err)
	})
}

func TestSet(t *testing.T) {
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)
	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] saving cache...", "key").Once()
		mockCacher.On("Set", "key", "data").Return(errDummy).Once()
		mockLogger.On("Error", "[%s] failed saving cache: %s", "key", errDummy.Error()).Once()
		c := newCacherLog(mockCacher, mockLogger)

		err := c.Set("key", "data")
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] saving cache...", "key").Once()
		mockCacher.On("Set", "key", "data").Return(nil).Once()
		mockLogger.On("Debug", "[%s] cache saved (%s)", "key", time.Second.Truncate(time.Microsecond)).Once()
		c := newCacherLog(mockCacher, mockLogger)

		err := c.Set("key", "data")
		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)
	t.Run("error", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] deleting cache...", "key").Once()
		mockCacher.On("Delete", "key").Return(errDummy).Once()
		mockLogger.On("Error", "[%s] failed deleting cache: %s", "key", errDummy.Error()).Once()
		c := newCacherLog(mockCacher, mockLogger)

		err := c.Delete("key")
		assert.Error(t, err)
		assert.EqualError(t, err, errDummy.Error())
	})

	t.Run("ok", func(t *testing.T) {
		mockLogger.On("Trace", "[%s] deleting cache...", "key").Once()
		mockCacher.On("Delete", "key").Return(nil).Once()
		mockLogger.On("Debug", "[%s] cache deleted (%s)", "key", time.Second.Truncate(time.Microsecond)).Once()
		c := newCacherLog(mockCacher, mockLogger)

		err := c.Delete("key")
		assert.NoError(t, err)
	})
}

func TestClose(t *testing.T) {
	mockCacher := new(mocks.Cacher)
	mockLogger := new(mocks.Logger)
	t.Run("ok", func(t *testing.T) {
		mockCacher.On("Close").Return(nil).Once()
		c := newCacherLog(mockCacher, mockLogger)

		err := c.Close()
		assert.NoError(t, err)
	})
}
