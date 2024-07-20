package palindromephrase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"

	assert.Equal(t, true, isPalindrome(s))

}
