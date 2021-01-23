package malscraper

import (
	e "errors"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/mal-plugin/cache/bigcache"
	"github.com/stretchr/testify/assert"
)

var errDummy = e.New("dummy error")
var mal *Malscraper
var sleepDur = 5 * time.Second

func init() {
	mal, _ = NewDefault()
}

func TestNew(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		createCache = func(time.Duration) (*bigcache.Client, error) {
			return nil, errDummy
		}
		m, err := New(Config{})
		assert.Nil(t, m)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrInitCache.Error())
	})

	t.Run("ok", func(t *testing.T) {
		createCache = bigcache.New
		m, err := New(Config{})
		assert.NotNil(t, m)
		assert.NoError(t, err)
	})
}

func TestDefault(t *testing.T) {
	m, err := NewDefault()
	assert.NotNil(t, m)
	assert.NoError(t, err)
}

func TestNewNoCache(t *testing.T) {
	m, err := NewNoCache()
	assert.NotNil(t, m)
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := &Malscraper{}
		err := m.Close()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		m, _ := NewNoCache()
		err := m.Close()
		assert.NoError(t, err)
	})
}
