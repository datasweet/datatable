package datatable_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestEmptyCopy(t *testing.T) {
	tb := New(t)
	cpy := tb.Copy(serie.EmptyCopy)

	assert.NotNil(t, cpy)
	assert.NotSame(t, tb, cpy)
	assert.Equal(t, 0, cpy.NumRows())
	assert.Equal(t, tb.NumCols(), cpy.NumCols())
}

func TestShallowCopy(t *testing.T) {
	tb := New(t)
	cpy := tb.Copy(serie.ShallowCopy)
	assert.NotNil(t, cpy)
	assert.NotSame(t, tb, cpy)
	assert.Equal(t, tb.NumRows(), cpy.NumRows())
	assert.Equal(t, tb.NumCols(), cpy.NumCols())

	checkTable(t, cpy,
		"champ", "champion", "win", "loose", "winRate", "sum", "ok",
		"Malzahar", "MALZAHAR", 10, 6, "62.5 %", 696.0, true,
		"Xerath", "XERATH", 20, 5, "80 %", 696.0, true,
		"Teemo", "TEEMO", 666, 666, "50 %", 696.0, true,
	)

	// edit original => should share
}
