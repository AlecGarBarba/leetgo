package keysrooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyRooms(t *testing.T) {

	rooms := [][]int{}
	assert.Equal(t, true, canVisitAllRooms(rooms), "Empty rooms should return true")

}

/*
Example 1:

Input: rooms = [[1],[2],[3],[]]
Output: true
*/
func TestRoomsWithAllKeys(t *testing.T) {

	rooms := [][]int{{1}, {2}, {3}, {}}
	assert.Equal(t, true, canVisitAllRooms(rooms), "all keys are available")

}

func TestRoomsWithMissingKey(t *testing.T) {

	rooms := [][]int{{1}, {3}, {2}, {}}
	assert.Equal(t, false, canVisitAllRooms(rooms), "Rooms have missing key")

}
