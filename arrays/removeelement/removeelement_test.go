package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	/*
			Input: nums = [3,2,2,3], val = 3
		Output: 2, nums = [2,2,_,_]
	*/
	nums := []int{3, 2, 2, 3}
	val := 3
	assert.Equal(t, 2, removeElement(nums, val), "Occurence count mismatch")
	assert.Equal(t, []int{2, 2}, nums[:2], "Remove element in place failed")

}

func TestExample2(t *testing.T) {
	/*
		input = [0,1,2,2,3,0,4,2]
		val = 2
		expected:
			[0,1,4,0,3]
	*/
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	assert.Equal(t, 5, removeElement(input, val), "Occurence count mismatch")
	assert.Equal(t, []int{0, 1, 3, 0, 4}, input[:5], "Remove element in place failed")
}
