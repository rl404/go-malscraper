package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStrToTime(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		test, valid := StrToTime("")
		assert.Equal(t, test, time.Time{})
		assert.False(t, valid)
	})

	t.Run("now", func(t *testing.T) {
		test, valid := StrToTime("now")
		assert.NotZero(t, test)
		assert.True(t, valid)
	})

	t.Run("seconds ago", func(t *testing.T) {
		now := time.Now()
		now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)
		test, valid := StrToTime("1 second ago")
		assert.Equal(t, test, now.Add(-time.Second))
		assert.True(t, valid)

		test, valid = StrToTime("5 seconds ago")
		assert.Equal(t, test, now.Add(-5*time.Second))
		assert.True(t, valid)
	})

	t.Run("minutes ago", func(t *testing.T) {
		now := time.Now()
		now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)
		test, valid := StrToTime("1 minute ago")
		assert.Equal(t, test, now.Add(-time.Minute))
		assert.True(t, valid)

		test, valid = StrToTime("5 minutes ago")
		assert.Equal(t, test, now.Add(-5*time.Minute))
		assert.True(t, valid)
	})

	t.Run("hours ago", func(t *testing.T) {
		now := time.Now()
		now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)
		test, valid := StrToTime("1 hour ago")
		assert.Equal(t, test, now.Add(-time.Hour))
		assert.True(t, valid)

		test, valid = StrToTime("5 hours ago")
		assert.Equal(t, test, now.Add(-5*time.Hour))
		assert.True(t, valid)
	})

	t.Run("Today, 8:48 AM", func(t *testing.T) {
		now := time.Now()
		now = time.Date(now.Year(), now.Month(), now.Day(), 8, 48, 0, 0, time.Local)
		test, valid := StrToTime("Today, 8:48 AM")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Yesterday, 8:48 AM", func(t *testing.T) {
		now := time.Now().Add(-24 * time.Hour)
		now = time.Date(now.Year(), now.Month(), now.Day(), 8, 48, 0, 0, time.Local)
		test, valid := StrToTime("Yesterday, 8:48 AM")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Apr 17, 8:57 PM", func(t *testing.T) {
		now := time.Date(time.Now().Year(), 4, 17, 20, 57, 0, 0, time.Local)
		test, valid := StrToTime("Apr 17, 8:57 PM")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Apr 17, 2017 8:57 PM", func(t *testing.T) {
		now := time.Date(2017, 4, 17, 20, 57, 0, 0, time.Local)
		test, valid := StrToTime("Apr 17, 2017 8:57 PM")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Apr 7, 2013, 01:58 (JST)", func(t *testing.T) {
		loc := time.FixedZone("JST", 0)
		now := time.Date(2013, 4, 7, 1, 58, 0, 0, loc)
		test, valid := StrToTime("Apr 7, 2013, 01:58 (JST)")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Apr 3, 1998", func(t *testing.T) {
		now := time.Date(1998, 4, 3, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("Apr 3, 1998")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Apr ??, 1998", func(t *testing.T) {
		now := time.Date(1998, 4, 1, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("Apr ??, 1998")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("??? ??, 1998", func(t *testing.T) {
		now := time.Date(1998, 1, 1, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("??? ??, 1998")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Jan, 2021", func(t *testing.T) {
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("Jan, 2021")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Jan 2021", func(t *testing.T) {
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("Jan 2021")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Jan 3,", func(t *testing.T) {
		now := time.Date(0, 1, 3, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("Jan 3,")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("Jan 3", func(t *testing.T) {
		now := time.Date(0, 1, 3, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("Jan 3")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("2021", func(t *testing.T) {
		now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("2021")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("10-25-02", func(t *testing.T) {
		now := time.Date(2002, 10, 25, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("10-25-02")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("10-??-02", func(t *testing.T) {
		now := time.Date(2002, 10, 1, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("10-??-02")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("??-??-02", func(t *testing.T) {
		now := time.Date(2002, 1, 1, 0, 0, 0, 0, time.Local)
		test, valid := StrToTime("??-??-02")
		assert.Equal(t, test, now)
		assert.True(t, valid)
	})

	t.Run("invalid", func(t *testing.T) {
		test, valid := StrToTime("random")
		assert.Equal(t, test, time.Time{})
		assert.False(t, valid)
	})
}

func TestStrToDate(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		y, m, d := StrToDate("")
		assert.Equal(t, y, 0)
		assert.Equal(t, m, 0)
		assert.Equal(t, d, 0)
	})

	t.Run("Apr 3, 1998", func(t *testing.T) {
		y, m, d := StrToDate("Apr 3, 1998")
		assert.Equal(t, y, 1998)
		assert.Equal(t, m, 4)
		assert.Equal(t, d, 3)
	})

	t.Run("Apr ??, 1998", func(t *testing.T) {
		y, m, d := StrToDate("Apr ??, 1998")
		assert.Equal(t, y, 1998)
		assert.Equal(t, m, 4)
		assert.Equal(t, d, 0)
	})

	t.Run("??? ??, 1998", func(t *testing.T) {
		y, m, d := StrToDate("??? ??, 1998")
		assert.Equal(t, y, 1998)
		assert.Equal(t, m, 0)
		assert.Equal(t, d, 0)
	})

	t.Run("Apr 1998", func(t *testing.T) {
		y, m, d := StrToDate("Apr 1998")
		assert.Equal(t, y, 1998)
		assert.Equal(t, m, 4)
		assert.Equal(t, d, 0)
	})

	t.Run("Apr 3", func(t *testing.T) {
		y, m, d := StrToDate("Apr 3")
		assert.Equal(t, y, 0)
		assert.Equal(t, m, 4)
		assert.Equal(t, d, 3)
	})

	t.Run("1998", func(t *testing.T) {
		y, m, d := StrToDate("1998")
		assert.Equal(t, y, 1998)
		assert.Equal(t, m, 0)
		assert.Equal(t, d, 0)
	})

	t.Run("Apr 9 2017 10:57 PM", func(t *testing.T) {
		y, m, d := StrToDate("Apr 9 2017 10:57 PM")
		assert.Equal(t, y, 2017)
		assert.Equal(t, m, 4)
		assert.Equal(t, d, 9)
	})

	t.Run("Apr 9 2017 10:57 (JST)", func(t *testing.T) {
		y, m, d := StrToDate("Apr 9 2017 10:57 (JST)")
		assert.Equal(t, y, 2017)
		assert.Equal(t, m, 4)
		assert.Equal(t, d, 9)
	})

	t.Run("10-25-02", func(t *testing.T) {
		y, m, d := StrToDate("10-25-02")
		assert.Equal(t, y, 2002)
		assert.Equal(t, m, 10)
		assert.Equal(t, d, 25)
	})

	t.Run("10-25-96", func(t *testing.T) {
		y, m, d := StrToDate("10-25-96")
		assert.Equal(t, y, 1996)
		assert.Equal(t, m, 10)
		assert.Equal(t, d, 25)
	})

	t.Run("10-??-02", func(t *testing.T) {
		y, m, d := StrToDate("10-??-02")
		assert.Equal(t, y, 2002)
		assert.Equal(t, m, 10)
		assert.Equal(t, d, 0)
	})

	t.Run("10-??-96", func(t *testing.T) {
		y, m, d := StrToDate("10-??-96")
		assert.Equal(t, y, 1996)
		assert.Equal(t, m, 10)
		assert.Equal(t, d, 0)
	})

	t.Run("??-??-02", func(t *testing.T) {
		y, m, d := StrToDate("??-??-02")
		assert.Equal(t, y, 2002)
		assert.Equal(t, m, 0)
		assert.Equal(t, d, 0)
	})

	t.Run("??-??-96", func(t *testing.T) {
		y, m, d := StrToDate("??-??-96")
		assert.Equal(t, y, 1996)
		assert.Equal(t, m, 0)
		assert.Equal(t, d, 0)
	})

	t.Run("invalid", func(t *testing.T) {
		y, m, d := StrToDate("random")
		assert.Equal(t, y, 0)
		assert.Equal(t, m, 0)
		assert.Equal(t, d, 0)
	})
}

func TestGetDuration(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		d := GetDuration("")
		assert.Zero(t, d)
	})

	t.Run("good", func(t *testing.T) {
		d := GetDuration("3 hr. 30 min. 20 sec.")
		assert.Equal(t, d, 12620)
	})
}

func TestSecondToString(t *testing.T) {
	d := SecondToString(10921)
	assert.Equal(t, d, "03:02:01")
}
