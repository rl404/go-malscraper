package malscraper

import (
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
)

// Config is malscraper configuration.
type Config struct {
	// Cache interface with basic get & set functions.
	// Can use your own cacher interface.
	Cacher internal.Cacher
	// Cache expired time. Will be used to initiating `Cacher`
	// using in-memory (bigcache) if `Cacher` is empty.
	CacheTime time.Duration

	// Does malscraper need to automatically clean any image and video url.
	// For more information, please read `ImageURLCleaner()` and `VideoURLCleaner()`
	// function in `pkg/utils/utils.go`.
	CleanImageURL bool
	CleanVideoURL bool

	// Log interface. Can use your own logger interface.
	Logger internal.Logger
	// Log Level. Show only error as default. Value should be chosen from constant.
	// Will be used to intiating `Logger` if `Logger` is empty.
	LogLevel int
	// Colorful log. Will be used to intiating `Logger` if `Logger` is empty.
	LogColor bool
}

func (c *Config) init() (err error) {
	if c.Logger == nil {
		c.Logger = internal.NewLogger(c.LogLevel, c.LogColor)
	}

	if c.Cacher == nil {
		c.Cacher, err = internal.NewCacher(c.Logger, c.CacheTime)
		if err != nil {
			c.Logger.Error("failed initiating cache: %s", err.Error())
			return errors.ErrInitCache
		}
	}

	return nil
}
