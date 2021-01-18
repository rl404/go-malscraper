package internal

import (
	"errors"
	"time"

	"github.com/rl404/mal-plugin/cache/bigcache"
)

// Cacher is caching interface required for malscraper
// caching system. If you use custom caching, try to
// implement this interface to your cacher.
type Cacher interface {
	// Get data from cache. The returned value will be
	// assigned to param `data`. Param `data` should
	// be a pointer just like when using json.Unmarshal.
	Get(key string, data interface{}) error
	// Save data to cache. Set and Get should be using
	// the same encoding method for example, json.Marshal
	// for Set and json.Unmarshal for Get.
	Set(key string, data interface{}) error
	// Delete data from cache.
	Delete(key string) error
	// Close cache connection.
	Close() error
}

// NewCacher to create new default cacher.
func NewCacher(l Logger, expiredTime time.Duration) (Cacher, error) {
	if expiredTime <= 0 {
		expiredTime = 24 * time.Hour
	}
	l.Info("caching: %v", expiredTime)
	return bigcache.New(expiredTime)
}

type noCache struct{}

// NewNoCacher to create cacher without actually caching.
func NewNoCacher() Cacher {
	return &noCache{}
}

// Set will just return nil.
func (c *noCache) Set(key string, data interface{}) error {
	return nil
}

// Get will just return error to simulate as if data is
// not in cache.
func (c *noCache) Get(key string, data interface{}) error {
	return errors.New("not using cache")
}

// Delete will just return nil.
func (c *noCache) Delete(key string) error {
	return nil
}

// Close will just return nil.
func (c *noCache) Close() error {
	return nil
}
