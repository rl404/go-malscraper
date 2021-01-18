package cacher

import (
	"time"

	"github.com/rl404/go-malscraper/internal"
)

// Cacher intercepts request to check the requested
// data to cache before actually access and parse
// MyAnimeList web.
type Cacher struct {
	api    internal.API
	cacher internal.Cacher
	logger internal.Logger
}

// New to create new cacher.
func New(api internal.API, c internal.Cacher, l internal.Logger) internal.API {
	return &Cacher{
		api:    api,
		cacher: newCacherLog(c, l),
		logger: l,
	}
}

// Simple cacher wrapper withh log to prevent writing
// repetitive log code.
type cacherLog struct {
	cacher internal.Cacher
	logger internal.Logger
}

// Testable time since func.
var timeSince = time.Since

func newCacherLog(c internal.Cacher, l internal.Logger) internal.Cacher {
	return &cacherLog{
		cacher: c,
		logger: l,
	}
}

// Get to get data from cache with log.
func (c cacherLog) Get(key string, data interface{}) error {
	c.logger.Trace("[%s] retrieving cache...", key)
	t := time.Now()
	if err := c.cacher.Get(key, data); err != nil {
		c.logger.Warn("[%s] failed retrieving cache: %s", key, err.Error())
		return err
	}
	c.logger.Debug("[%s] cache found (%s)", key, timeSince(t).Truncate(time.Microsecond))
	return nil
}

// Set to save data to cache with log.
func (c cacherLog) Set(key string, data interface{}) error {
	c.logger.Trace("[%s] saving cache...", key)
	t := time.Now()
	if err := c.cacher.Set(key, data); err != nil {
		c.logger.Error("[%s] failed saving cache: %s", key, err.Error())
		return err
	}
	c.logger.Debug("[%s] cache saved (%s)", key, timeSince(t).Truncate(time.Microsecond))
	return nil
}

// Delete to delete data in cache with log.
func (c cacherLog) Delete(key string) error {
	c.logger.Trace("[%s] deleting cache...", key)
	t := time.Now()
	if err := c.cacher.Delete(key); err != nil {
		c.logger.Error("[%s] failed deleting cache: %s", key, err.Error())
		return err
	}
	c.logger.Debug("[%s] cache deleted (%s)", key, timeSince(t).Truncate(time.Microsecond))
	return nil
}

// Close to close cache connection.
func (c cacherLog) Close() error {
	return c.cacher.Close()
}