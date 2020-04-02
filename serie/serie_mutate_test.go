package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrowZero(t *testing.T) {
	s := NewSerieInt(t)
	assert.Equal(t, 100, s.Len())
	assert.NoError(t, s.Error())
	s.Grow(10)
	assert.NoError(t, s.Error())
	assert.Equal(t, 110, s.Len())
	assert.Equal(t, []interface{}{
		31, 23, 98, 3, 59, 67, 5, 5, 87, 18,
		3, 88, 7, 63, 29, 62, 37, 66, 87, 26,
		24, 5, 62, 75, 69, 56, 15, 59, 40, 34,
		68, 32, 34, 29, 90, 21, 8, 8, 100, 64,
		30, 56, 73, 2, 65, 74, 3, 26, 92, 46,
		6, 100, 35, 17, 91, 55, 99, 87, 9, 25,
		55, 76, 39, 78, 43, 99, 35, 90, 36, 27,
		52, 65, 33, 49, 84, 87, 42, 92, 27, 65,
		48, 47, 74, 98, 76, 88, 18, 100, 69, 57,
		69, 90, 74, 25, 64, 37, 63, 61, 85, 12,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, s.Values())
}

func TestShrink(t *testing.T) {
	s := NewSerieInt(t)
	assert.Equal(t, 100, s.Len())
	assert.NoError(t, s.Error())
	s.Shrink(90)
	assert.NoError(t, s.Error())
	assert.Equal(t, 10, s.Len())
	assert.Equal(t, []interface{}{31, 23, 98, 3, 59, 67, 5, 5, 87, 18}, s.Values())

	s.Shrink(25)
	assert.Error(t, s.Error())
	assert.Equal(t, 10, s.Len())
	assert.Equal(t, []interface{}{31, 23, 98, 3, 59, 67, 5, 5, 87, 18}, s.Values())

	// s.Shrink(10)
	// assert.NoError(t, s.Error())
	// assert.Equal(t, 0, s.Len())
	// assert.Equal(t, []interface{}{}, s.Values())
}
