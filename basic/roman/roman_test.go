package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleDigitRoman(t *testing.T) {

	assert.EqualValues(t, 1, romanToInt("I"))

	assert.EqualValues(t, 3, romanToInt("III"))

	assert.EqualValues(t, 4, romanToInt("IV"))

	assert.EqualValues(t, 5, romanToInt("V"))

	assert.EqualValues(t, 6, romanToInt("VI"))

	assert.EqualValues(t, 9, romanToInt("IX"))
	assert.EqualValues(t, 10, romanToInt("X"))

}

func TestTeens(t *testing.T) {
	assert.EqualValues(t, 11, romanToInt("XI"))
	assert.EqualValues(t, 14, romanToInt("XIV"))
	assert.EqualValues(t, 18, romanToInt("XVIII"))
	assert.EqualValues(t, 19, romanToInt("XIX"))
}

func TestDoubleDigits(t *testing.T) {
	assert.EqualValues(t, 40, romanToInt("XL"))
	assert.EqualValues(t, 44, romanToInt("XLIV"))
	assert.EqualValues(t, 50, romanToInt("L"))
	assert.EqualValues(t, 90, romanToInt("XC"))
}

func TestTripleDigits(t *testing.T) {
	assert.EqualValues(t, 100, romanToInt("C"))
	assert.EqualValues(t, 400, romanToInt("CD"))
	assert.EqualValues(t, 500, romanToInt("D"))
	assert.EqualValues(t, 900, romanToInt("CM"))
}

func TestThousands(t *testing.T) {
	assert.EqualValues(t, 1000, romanToInt("M"))
	assert.EqualValues(t, 3999, romanToInt("MMMCMXCIX"))
}
