package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToNum(t *testing.T) {
	assert.Equal(t, 1234567, StrToNum(" 1,234,567 "))
}

func TestStrToFloat(t *testing.T) {
	assert.Equal(t, 1234567.89, StrToFloat(" 1,234,567.89 "))
}

func TestGetValueFromSplit(t *testing.T) {
	t.Run("out of index", func(t *testing.T) {
		str := "https://myanimelist.net/anime/2/abc"
		assert.Equal(t, "", GetValueFromSplit(str, "/", 10))
	})

	t.Run("ok", func(t *testing.T) {
		str := "https://myanimelist.net/anime/2/abc"
		assert.Equal(t, "2", GetValueFromSplit(str, "/", 4))
	})
}

func ExampleGetValueFromSplit() {
	url := "https://myanimelist.net/anime/2/abc"
	fmt.Println(GetValueFromSplit(url, "/", 4))
	fmt.Println(GetValueFromSplit(url, "/", 3))
	fmt.Println(GetValueFromSplit(url, "/", 10)) // return empty string
	// Output:
	// 2
	// anime
	//
}

func TestThousand(t *testing.T) {
	assert.Equal(t, "-1,000", Thousand(-1000))
	assert.Equal(t, "-1", Thousand(-1))
	assert.Equal(t, "1", Thousand(1))
	assert.Equal(t, "1,000", Thousand(1000))
	assert.Equal(t, "1,234,567", Thousand(1234567))
	assert.Equal(t, "1,234,567,890", Thousand(1234567890))
}

func TestEllipsis(t *testing.T) {
	assert.Equal(t, "123...", Ellipsis("123456789", 3))
	assert.Equal(t, "123456789", Ellipsis("123456789", 20))
}
