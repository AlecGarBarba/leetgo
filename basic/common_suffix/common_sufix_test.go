package commonsuffix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyArray(t *testing.T) {
	assert.Equal(t, "", longestCommonPrefix([]string{"", ""}))
}

func TestNoCommonPrefix(t *testing.T) {
	assert.Equal(t, "", longestCommonPrefix([]string{"abc", "def"}))
}

func TestOneCommonPrefix(t *testing.T) {
	assert.Equal(t, "a", longestCommonPrefix([]string{"abc", "ade"}))
}

func TestOneElement(t *testing.T) {
	assert.Equal(t, "word", longestCommonPrefix([]string{"word"}))
}

func TestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
