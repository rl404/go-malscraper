package cacher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKey(t *testing.T) {
	t.Run("without-param", func(t *testing.T) {
		key := getKey(keyProducers)
		assert.Equal(t, key, "mal:producers")
	})

	t.Run("with-param", func(t *testing.T) {
		key := getKey(keyProducer, 1, 2)
		assert.Equal(t, key, "mal:producer:1:2")
	})
}
