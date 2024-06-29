package lettercombinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBacktrackEmpty(t *testing.T) {
	assert.Equal(t, []string{}, letterCombinationsBackTrack(""))
}

func TestBacktrackSingleDigit(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinationsBackTrack("2"))
	assert.Equal(t, []string{"d", "e", "f"}, letterCombinationsBackTrack("3"))
	assert.Equal(t, []string{"g", "h", "i"}, letterCombinationsBackTrack("4"))
	assert.Equal(t, []string{"j", "k", "l"}, letterCombinationsBackTrack("5"))
	assert.Equal(t, []string{"m", "n", "o"}, letterCombinationsBackTrack("6"))
	assert.Equal(t, []string{"p", "q", "r", "s"}, letterCombinationsBackTrack("7"))
	assert.Equal(t, []string{"t", "u", "v"}, letterCombinationsBackTrack("8"))
	assert.Equal(t, []string{"w", "x", "y", "z"}, letterCombinationsBackTrack("9"))
}

func TestBacktrackDoubleDigit(t *testing.T) {
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinationsBackTrack("23"))
	assert.Equal(t, []string{"ag", "ah", "ai", "bg", "bh", "bi", "cg", "ch", "ci"}, letterCombinationsBackTrack("24"))
}

func TestBacktrackTripleDigit(t *testing.T) {
	assert.Equal(t, []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, letterCombinationsBackTrack("234"))
}
