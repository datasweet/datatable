package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestHead(t *testing.T) {
	s := serie.NewInt64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Head(5), "1 2 3 4 5")
	assertSerieEq(t, s.Head(1), "1")
	assertSerieEq(t, s.Head(9), "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Head(10), "nil")
	assertSerieEq(t, s.Head(0), "nil")
	assertSerieEq(t, s.Head(-1), "nil")
}

func TestTail(t *testing.T) {
	s := serie.NewInt64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Tail(5), "5 6 7 8 9")
	assertSerieEq(t, s.Tail(1), "9")
	assertSerieEq(t, s.Tail(9), "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Tail(10), "nil")
	assertSerieEq(t, s.Tail(0), "nil")
	assertSerieEq(t, s.Tail(-1), "nil")
}

func TestSubset(t *testing.T) {
	s := serie.NewInt64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Subset(5, 3), "5 6 7")
	assertSerieEq(t, s.Subset(8, 2), "8 9")
	assertSerieEq(t, s.Subset(8, 3), "nil")
	assertSerieEq(t, s.Subset(9, 1), "9")
	assertSerieEq(t, s.Subset(0, 5), "nil")
	assertSerieEq(t, s.Subset(1, 1), "1")
	assertSerieEq(t, s.Subset(5, 0), "nil")
	assertSerieEq(t, s.Subset(5, -1), "nil")
	assertSerieEq(t, s.Subset(-1, 5), "nil")
	assertSerieEq(t, s.Subset(10, 5), "nil")
}

func TestPick(t *testing.T) {
	s := serie.NewInt64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Pick(1, 5, -1, 9, 10, 0), "1 5 #NULL! 9 #NULL! #NULL!")
	assertSerieEq(t, s.Pick(), "nil")
	assertSerieEq(t, s.Pick(1), "1")
}

func TestFindRows(t *testing.T) {
	s := serie.NewInt64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assert.Equal(t, []int{1, 3, 5, 7, 9}, s.FindRows(func(val serie.Value) bool {
		i := val.Val().(int64)
		return i%2 == 1
	}))
}

func TestFilter(t *testing.T) {
	s := serie.NewInt64(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, "1 2 3 4 5 6 7 8 9")
	assertSerieEq(t, s.Filter(func(val serie.Value) bool {
		i := val.Val().(int64)
		return i%2 == 1
	}), "1 3 5 7 9")
}
