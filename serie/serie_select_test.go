package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/datatable/value"
)

func TestHead(t *testing.T) {
	s := serie.Int64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Head(5), "1 2 3 4 5")
	assertSerieEq(t, s.Head(1), "1")
	assertSerieEq(t, s.Head(9), "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Head(10), "nil")
	assertSerieEq(t, s.Head(0), "nil")
	assertSerieEq(t, s.Head(-1), "nil")
}

func TestTail(t *testing.T) {
	s := serie.Int64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Tail(5), "5 6 7 8 9")
	assertSerieEq(t, s.Tail(1), "9")
	assertSerieEq(t, s.Tail(9), "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Tail(10), "nil")
	assertSerieEq(t, s.Tail(0), "nil")
	assertSerieEq(t, s.Tail(-1), "nil")
}

func TestSubset(t *testing.T) {
	s := serie.Int64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Subset(4, 3), "5 6 7")
	assertSerieEq(t, s.Subset(7, 2), "8 9")
	assertSerieEq(t, s.Subset(7, 3), "nil")
	assertSerieEq(t, s.Subset(8, 1), "9")
	assertSerieEq(t, s.Subset(0, 5), "1 2 3 4 5")
	assertSerieEq(t, s.Subset(0, 1), "1")
	assertSerieEq(t, s.Subset(5, 0), "nil")
	assertSerieEq(t, s.Subset(5, -1), "nil")
	assertSerieEq(t, s.Subset(-1, 5), "nil")
	assertSerieEq(t, s.Subset(10, 5), "nil")
}

func TestPick(t *testing.T) {
	s := serie.Int64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Pick(0, 4, -1, 8, 9, 1), "1 5 #NULL! 9 #NULL! 2")
	assertSerieEq(t, s.Pick(), "nil")
	assertSerieEq(t, s.Pick(0), "1")
}

func TestFindRows(t *testing.T) {
	s := serie.Int64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assert.Equal(t, []int{0, 2, 4, 6, 8}, s.FindRows(func(val value.Value) bool {
		i := val.Val().(int64)
		return i%2 == 1
	}))
}

func TestFilter(t *testing.T) {
	s := serie.Int64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Filter(func(val value.Value) bool {
		i := val.Val().(int64)
		return i%2 == 1
	}), "1 3 5 7 9")
}

func TestDistinct(t *testing.T) {
	s := NewSerieInt(t)
	assertSerieIntEq(t, s.Distinct(),
		31, 23, 98, 3, 59, 67, 5, 87, 18, 88,
		7, 63, 29, 62, 37, 66, 26, 24, 75, 69,
		56, 15, 40, 34, 68, 32, 90, 21, 8, 100,
		64, 30, 73, 2, 65, 74, 92, 46, 6, 35,
		17, 91, 55, 99, 9, 25, 76, 39, 78, 43,
		36, 27, 52, 33, 49, 84, 42, 48, 47, 57,
		61, 85, 12,
	)
}
