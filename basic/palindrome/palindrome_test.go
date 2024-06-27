package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeFalse(t *testing.T) {
	assert.False(t, isPalindrome(10))

}

func TestPalindromeTrue(t *testing.T) {
	assert.True(t, isPalindrome(121))
}
