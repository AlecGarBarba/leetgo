package maxarea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestExample1(t *testing.T){
	/*
		Input: height = [1,8,6,2,5,4,8,3,7]
		Output: 49
	*/
	assert.Equal(t, 49, maxArea([]int{1,8,6,2,5,4,8,3,7}), "Output should be 49")
}