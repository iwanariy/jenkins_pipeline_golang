package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fizzbuzz_1(t *testing.T) {
	assert.Equal(t, fizzbuzz("1"), "1", "")
	assert.Equal(t, fizzbuzz("3"), "fizz", "")
	assert.Equal(t, fizzbuzz("5"), "buzz", "")
	assert.Equal(t, fizzbuzz("15"), "fizzbuzz", "")
}
