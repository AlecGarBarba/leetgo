package keysrooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Example 1:

Input: rooms = [[1],[2],[3],[]]
Output: true
*/
func TestRoomsWithAllKeys(t *testing.T) {

	rooms := [][]int{{1}, {2}, {3}, {}}
	assert.Equal(t, true, canVisitAllRooms(rooms), "all keys are available")

}
