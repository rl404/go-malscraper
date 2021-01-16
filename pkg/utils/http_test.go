package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	assert.Equal(t, "http://a.com", BuildURL("http://a.com"))
	assert.Equal(t, "http://a.com/b/c", BuildURL("http://a.com", "b", "c"))
}

func TestBuildURLWithQuery(t *testing.T) {
	assert.Equal(t, "http://a.com", BuildURLWithQuery(nil, "http://a.com"))
	assert.Equal(t, "http://a.com/b/c?d=1&e=2", BuildURLWithQuery(map[string]interface{}{
		"d": 1,
		"e": "2",
	}, "http://a.com", "b", "c"))
}
