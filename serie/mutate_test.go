package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	s := serie.Int()
	assertSerieEq(t, s)

	s.Append(1, 2, 3, 4, "5")
	assertSerieEq(t, s, 1, 2, 3, 4, 5)

	s.Append(nil, 6, 7, "8", 9, 10)
	assertSerieEq(t, s, 1, 2, 3, 4, 5, 0, 6, 7, 8, 9, 10)
}

func TestPrepend(t *testing.T) {
	s := serie.Int()
	assertSerieEq(t, s)

	assert.NoError(t, s.Prepend(1, 2, 3, 4, 5))
	assertSerieEq(t, s, 1, 2, 3, 4, 5)

	assert.NoError(t, s.Prepend(-4, -3, -2, -1, 0))
	assertSerieEq(t, s, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5)
}

func TestInsert(t *testing.T) {
	s := serie.Int(1, 2, 3, 4, 5)
	assertSerieEq(t, s, 1, 2, 3, 4, 5)

	assert.NoError(t, s.Insert(2, 7, 8, 9, 10))
	assertSerieEq(t, s, 1, 2, 7, 8, 9, 10, 3, 4, 5)

	assert.Error(t, s.Insert(-1, 7, 8, 9, 10))
	assert.Error(t, s.Insert(101, 7, 8, 9, 10))
}

func TestSet(t *testing.T) {
	s := serie.Int(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	assert.Error(t, s.Set(-1, 100))
	assert.Error(t, s.Set(10, 100))
	assert.Error(t, s.Set(0, []int{0, 1, 2, 4}))

	assert.NoError(t, s.Set(5, 555))
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 555, 6, 7, 8, 9)

	assert.NoError(t, s.Set(9, 999))
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 555, 6, 7, 8, 999)

	assert.NoError(t, s.Set(0, -5))
	assertSerieEq(t, s, -5, 1, 2, 3, 4, 555, 6, 7, 8, 999)

}

func TestDelete(t *testing.T) {
	s := serie.Int(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	assert.Error(t, s.Delete(-1))
	assert.Error(t, s.Delete(10))

	assert.NoError(t, s.Delete(5))
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 6, 7, 8, 9)

	assert.NoError(t, s.Delete(8))
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 6, 7, 8)

	assert.NoError(t, s.Delete(0))
	assertSerieEq(t, s, 1, 2, 3, 4, 6, 7, 8)
}

func TestGrow(t *testing.T) {
	s := serie.Int(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	assert.Error(t, s.Grow(-5))

	assert.NoError(t, s.Grow(5))
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0, 0)

	s = serie.IntN(0, 1, 2, 3, 4, 5)
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5)
	assert.NoError(t, s.Grow(5))
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, nil, nil, nil, nil, nil)
}

func TestShrink(t *testing.T) {
	s := serie.Int(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	assert.Error(t, s.Shrink(-5))
	assert.Error(t, s.Shrink(11))

	assert.NoError(t, s.Shrink(5))
	assertSerieEq(t, s, 0, 1, 2, 3, 4)

	assert.NoError(t, s.Shrink(5))
	assertSerieEq(t, s)
}

func TestConcat(t *testing.T) {
	s := serie.Int(0, 1, 2, 3, 4)
	assert.Error(t, s.Concat(serie.IntN(-1, -2, nil)))
	assert.NoError(t, s.Concat(serie.Int(6, 7, 8, 9, 10)))
	assertSerieEq(t, s, 0, 1, 2, 3, 4, 6, 7, 8, 9, 10)

	s = serie.StringN("Léon", "Marie", "Sophie", "Marcel")
	assertSerieEq(t, s, "Léon", "Marie", "Sophie", "Marcel")
	assert.NoError(t, s.Concat(serie.StringN("Marion", "Paul", "Marie", "Marcel")))
	assertSerieEq(t, s, "Léon", "Marie", "Sophie", "Marcel", "Marion", "Paul", "Marie", "Marcel")
}

func TestClear(t *testing.T) {
	s := serie.Int(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.Equal(t, 10, s.Len())
	s.Clear()
	assert.Equal(t, 0, s.Len())
}
