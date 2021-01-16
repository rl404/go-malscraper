package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPercent(t *testing.T) {
	t.Run("divide by 0", func(t *testing.T) {
		assert.Equal(t, 0.0, GetPercent(10, 0))
	})

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 333.33, GetPercent(10, 3))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, 333.33, GetPercent(10.0, 3.0))
	})

	t.Run("digit", func(t *testing.T) {
		assert.Equal(t, 333.3, GetPercent(10.0, 3.0, 1))
	})
}

func ExampleGetPercent() {
	fmt.Println(GetPercent(10, 0))
	fmt.Println(GetPercent(10, 3))
	fmt.Println(GetPercent(10.0, 3.0))
	fmt.Println(GetPercent(10.0, 3.0, 1))
	// Output:
	// 0
	// 333.33
	// 333.33
	// 333.3
}
