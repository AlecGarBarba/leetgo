package wordsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMatrix(t *testing.T) {
	assert.False(t, exist([][]byte{}, "AB"))

}

func TestFindFirstWord(t *testing.T) {
	assert.True(t, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "A"))
	assert.False(t, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "Z"))
}

func TestExisting(t *testing.T) {
	assert.True(t, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}

func TestExisting2(t *testing.T) {
	assert.True(t, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
}
