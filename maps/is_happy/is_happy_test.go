package ishappy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.True(t, isHappy(19), "19 is a happy number")
}

func TestNonHappyNumber(t *testing.T) {
	assert.False(t, isHappy(2), "2 is not a happy number")
}
