package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayFilter(t *testing.T) {
	t.Run("contain empty", func(t *testing.T) {
		test := []string{"1", "2", "", "3"}
		result := []string{"1", "2", "3"}
		assert.True(t, reflect.DeepEqual(result, ArrayFilter(test)))
	})

	t.Run("empty result", func(t *testing.T) {
		test := []string{""}
		result := []string{}
		assert.True(t, reflect.DeepEqual(result, ArrayFilter(test)))
	})

	t.Run("nil", func(t *testing.T) {
		var test, result []string
		assert.True(t, reflect.DeepEqual(result, ArrayFilter(test)))
	})
}

func TestInArrayStr(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		arr := []string{"1", "2", "", "3"}
		assert.True(t, InArrayStr(arr, "2"))
	})

	t.Run("false", func(t *testing.T) {
		arr := []string{"1", "2", "", "3"}
		assert.False(t, InArrayStr(arr, "10"))
	})
}

func TestInArrayInt(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		arr := []int{1, 2, 3, 0}
		assert.True(t, InArrayInt(arr, 2))
	})

	t.Run("false", func(t *testing.T) {
		arr := []int{1, 2, 3, 0}
		assert.False(t, InArrayInt(arr, 10))
	})
}

func TestGetMapKey(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		m, v := map[int]string{0: "0", 1: "1", 2: "2", 3: "3"}, "2"
		k, f := GetMapKey(m, v)
		assert.Equal(t, k, 2)
		assert.True(t, f)
	})

	t.Run("false", func(t *testing.T) {
		m, v := map[int]string{0: "0", 1: "1", 2: "2", 3: "3"}, "20"
		_, f := GetMapKey(m, v)
		assert.False(t, f)
	})
}

func ExampleGetMapKey() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	fmt.Println(GetMapKey(m, "b"))
	fmt.Println(GetMapKey(m, "d"))
	// Output:
	// 2 true
	// 0 false
}

func TestUniqueInt(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var test, result []int
		assert.True(t, reflect.DeepEqual(result, UniqueInt(test)))
	})

	t.Run("empty array", func(t *testing.T) {
		test, result := []int{}, []int{}
		assert.True(t, reflect.DeepEqual(result, UniqueInt(test)))
	})

	t.Run("contain duplicate", func(t *testing.T) {
		test := []int{1, 2, 3, 4, 2}
		result := []int{1, 2, 3, 4}
		assert.True(t, reflect.DeepEqual(result, UniqueInt(test)))
	})
}
