package isanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {

	assert.True(t, isAnagram("anagram", "nagaram"))
}
func TestExample2(t *testing.T) {
	assert.False(t, isAnagram("rat", "car"))
}

func TestExample3(t *testing.T) {
	assert.False(t, isAnagram("a", "ab"))
}

func TestExample4(t *testing.T) {
	assert.False(t, isAnagram("a", "b"))
}

func TestExample5(t *testing.T) {
	assert.False(t, isAnagram("ab", "a"))
}
