package isisomorphic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:
*/
func TestExample1(t *testing.T) {

	assert.True(t, isIsomorphic("egg", "add"), "egg can map to add")
}

/*
Input: s = "foo", t = "bar"
Output: false
*/
func TestEXample2(t *testing.T) {
	assert.False(t, isIsomorphic("foo", "bar"), "foo cannot map to bar")
}

/*
Example 3:

Input: s = "paper", t = "title"
Output: true
*/
func TestExample3(t *testing.T) {
	assert.True(t, isIsomorphic("paper", "title"), "paper can map to title")
}

/*
example 4:
s =
"badc"
t =
"baba"
*/
func TestExample4(t *testing.T) {
	assert.False(t, isIsomorphic("badc", "baba"), "badc cannot map to baba")
}
