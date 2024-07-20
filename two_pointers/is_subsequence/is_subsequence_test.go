package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	/*
		Input: s = "abc", t = "ahbgdc"
		Output: true
	*/

	assert.True(t, isSubsequence("abc", "ahbgdc"), "Output should be true")

}

func TestFalse(t *testing.T) {
	assert.False(t, isSubsequence("b", "c"))
}

func TestEmpty(t *testing.T) {
	assert.True(t, isSubsequence("", ""))
}
