package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {

	int1 := []int{1, 2, 3, 0, 0, 0}
	merge(int1, 3, []int{2, 5, 6}, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, int1)

}

func TestExample2(t *testing.T) {

	int1 := []int{1}
	merge(int1, 1, []int{}, 0)
	assert.Equal(t, []int{1}, int1)

}

func TestExample3(t *testing.T) {
	int1 := []int{0}
	merge(int1, 0, []int{1}, 1)
	assert.Equal(t, []int{1}, int1)

}
