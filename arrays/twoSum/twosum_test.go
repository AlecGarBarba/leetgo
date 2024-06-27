package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidTwoSums(t *testing.T) {
	// GIVEN
	slice := []int{1, 2, 4}
	target := 3

	res := twoSum(slice, target)

	assert.ObjectsAreEqual(res, []int{0, 1})

	slice = []int{3, 2, 4}
	target = 6

	assert.ObjectsAreEqual(twoSum(slice, target), []int{1, 2})

}
