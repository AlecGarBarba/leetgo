package wordpattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	pattern := "abba"
	s := "dog cat cat dog"
	// Output: true
	assert.True(t, wordPattern(pattern, s), "Test case 1 failed")
}

func TestExample2(t *testing.T) {
	pattern := "abba"
	s := "dog cat cat fish"
	// Output: false
	assert.False(t, wordPattern(pattern, s), "Test case 2 failed")
}

func TestExample3(t *testing.T) {
	pattern := "aaaa"
	s := "dog cat cat dog"
	// Output: false
	assert.False(t, wordPattern(pattern, s), "Test case 3 failed")
}

func TestExample4(t *testing.T) {
	pattern := "abba"

	s := "dog dog dog dog"

	assert.False(t, wordPattern(pattern, s))
}
