package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	/*
		Input: nums = [3,2,3]
		Output: 3

	*/
	nums := []int{3, 2, 3}
	assert.Equal(t, 3, majorityElement(nums), "Majority element mismatch")
}

/*
Input: nums = [2,2,1,1,1,2,2]
Output: 2
*/
func TestExample2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	assert.Equal(t, 2, majorityElement(nums), "Majority element mismatch")
}

func TestExample3(t *testing.T) {
	/*
		Input: [3,3,4]
		Output: 3
	*/
	nums := []int{3, 3, 4}
	assert.Equal(t, 3, majorityElement(nums), "Majority element mismatch")
}
