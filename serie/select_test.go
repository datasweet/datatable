package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestHead(t *testing.T) {
	s := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Head(5), 1, 2, 3, 4, 5)
	assertSerieEq(t, s.Head(1), 1)
	assertSerieEq(t, s.Head(9), 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Head(10))
	assertSerieEq(t, s.Head(0))
	assertSerieEq(t, s.Head(-1))
	assertSerieEq(t, s.Head(5).Head(1), 1)
}

func TestTail(t *testing.T) {
	s := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Tail(5), 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Tail(1), 9)
	assertSerieEq(t, s.Tail(9), 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Tail(10))
	assertSerieEq(t, s.Tail(0))
	assertSerieEq(t, s.Tail(-1))
	assertSerieEq(t, s.Tail(5).Tail(1), 9)
}

func TestSubset(t *testing.T) {
	s := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Subset(4, 3), 5, 6, 7)
	assertSerieEq(t, s.Subset(7, 2), 8, 9)
	assertSerieEq(t, s.Subset(7, 3))
	assertSerieEq(t, s.Subset(8, 1), 9)
	assertSerieEq(t, s.Subset(0, 5), 1, 2, 3, 4, 5)
	assertSerieEq(t, s.Subset(0, 1), 1)
	assertSerieEq(t, s.Subset(5, 0))
	assertSerieEq(t, s.Subset(5, -1))
	assertSerieEq(t, s.Subset(-1, 5))
	assertSerieEq(t, s.Subset(10, 5))
}

func TestFilter(t *testing.T) {
	s := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	res, err := s.Filter(nil)
	assert.Error(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 0, res.Len())

	res, err = s.Filter(func(val float32) bool {
		return val == 3.14
	})
	assert.Error(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 0, res.Len())

	res, err = s.Filter(func(val int) bool {
		return val%2 == 1
	})
	assert.NoError(t, err)
	assertSerieEq(t, res, 1, 3, 5, 7, 9)
}

func TestDistinct(t *testing.T) {
	s := serie.Int(
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
	)
	assertSerieEq(t, s.Distinct(),
		31, 23, 98, 3, 59, 67, 5, 87, 18, 88,
		7, 63, 29, 62, 37, 66, 26, 24, 75, 69,
		56, 15, 40, 34, 68, 32, 90, 21, 8, 100,
		64, 30, 73, 2, 65, 74, 92, 46, 6, 35,
		17, 91, 55, 99, 9, 25, 76, 39, 78, 43,
		36, 27, 52, 33, 49, 84, 42, 48, 47, 57,
		61, 85, 12,
	)
}

func TestPick(t *testing.T) {
	s := serie.Int(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Pick(4, 3), 4, 3)
	assertSerieEq(t, s.Pick(-1), 0)
	assertSerieEq(t, s.Pick(0, -1, 4, 3, 9, 10), 0, 0, 4, 3, 9, 0)

	s = serie.IntN(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s.Pick(4, 3), 4, 3)
	assertSerieEq(t, s.Pick(-1), nil)
	assertSerieEq(t, s.Pick(0, -1, 4, 3, 9, 10), 0, nil, 4, 3, 9, nil)

}
