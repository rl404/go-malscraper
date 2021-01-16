package internal

import "github.com/rl404/go-malscraper/pkg/mallogger"

// Logger is logging interface for malscraper.
// If you use custom logger, try to
// implement this interface to your logger.
type Logger interface {
	Trace(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

// NewLogger to create new default logger.
func NewLogger(level int, color bool) Logger {
	if level == 0 {
		level = mallogger.LevelHigh
	}
	return mallogger.New(level, color)
}
