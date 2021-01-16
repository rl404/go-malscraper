package parser

import (
	"errors"
	"time"

	"github.com/rl404/go-malscraper/internal"
)

var errDummy = errors.New("dummy error")
var log = internal.NewLogger(0, false)
var sleepDur = 5 * time.Second
