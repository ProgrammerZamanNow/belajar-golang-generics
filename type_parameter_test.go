package belajar_golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Eko")
	assert.Equal(t, "Eko", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
