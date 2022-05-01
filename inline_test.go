package belajar_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 200))
	assert.Equal(t, int64(100), FindMin[int64](int64(100), int64(200)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Eko", "Kurniawan", "Khannedy",
	}

	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Eko", first)
}
