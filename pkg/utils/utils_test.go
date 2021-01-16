package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestURLCleaner(t *testing.T) {
	t.Run("no need", func(t *testing.T) {
		test := URLCleaner("abc", "img", false)
		assert.Equal(t, test, "abc")
	})

	t.Run("empty", func(t *testing.T) {
		test := URLCleaner("https://cdn.myanimelist.net/r/76x120/images/questionmark_50.gif?s=8e0400788aa6af2a2f569649493e2b0f", "img", true)
		assert.Equal(t, test, "")
	})

	t.Run("image", func(t *testing.T) {
		test := URLCleaner("https://cdn.myanimelist.net/r/80x120/images/manga/3/214566.jpg?s=48212bcd0396d503a01166149a29c67e", "img", true)
		assert.Equal(t, test, "https://cdn.myanimelist.net/images/manga/3/214566.jpg")
	})

	t.Run("video", func(t *testing.T) {
		test := URLCleaner("https://www.youtube.com/embed/qig4KOK2R2g?enablejsapi=1&wmode=opaque&autoplay=1", "video", true)
		assert.Equal(t, test, "https://www.youtube.com/watch?v=qig4KOK2R2g")
	})

	t.Run("random", func(t *testing.T) {
		test := URLCleaner("abc", "random", true)
		assert.Equal(t, test, "abc")
	})
}

func TestGetCurrentSeason(t *testing.T) {
	s := GetCurrentSeason()
	currentMonth := int(time.Now().Month())
	assert.Equal(t, s, GetSeasonName(currentMonth))
}

func TestGetSeasonName(t *testing.T) {
	mSeason := map[int]string{
		1:  "winter",
		2:  "winter",
		3:  "winter",
		4:  "spring",
		5:  "spring",
		6:  "spring",
		7:  "summer",
		8:  "summer",
		9:  "summer",
		10: "fall",
		11: "fall",
		12: "fall",
		13: "",
	}

	for m, s := range mSeason {
		assert.Equal(t, s, GetSeasonName(m))
	}
}
