package lettercombinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.Equal(t, []string{}, letterCombinations(""))
}

func TestSingleDigit(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
	assert.Equal(t, []string{"d", "e", "f"}, letterCombinations("3"))
	assert.Equal(t, []string{"g", "h", "i"}, letterCombinations("4"))
	assert.Equal(t, []string{"j", "k", "l"}, letterCombinations("5"))
	assert.Equal(t, []string{"m", "n", "o"}, letterCombinations("6"))
	assert.Equal(t, []string{"p", "q", "r", "s"}, letterCombinations("7"))
	assert.Equal(t, []string{"t", "u", "v"}, letterCombinations("8"))
	assert.Equal(t, []string{"w", "x", "y", "z"}, letterCombinations("9"))
}

func TestDoubleDigit(t *testing.T) {
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.Equal(t, []string{"ag", "ah", "ai", "bg", "bh", "bi", "cg", "ch", "ci"}, letterCombinations("24"))
}

func TestTripleDigit(t *testing.T) {
	assert.Equal(t, []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, letterCombinations("234"))
}
