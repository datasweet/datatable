package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	original := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	cpy := original.Copy()
	assert.NotSame(t, original, cpy)
	assert.Equal(t, original.Type(), cpy.Type())
	assert.Equal(t, original.Len(), cpy.Len())
	assertSerieEq(t, cpy, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	original.Set(4, 50)
	assertSerieEq(t, original, 1, 2, 3, 4, 50, 6, 7, 8, 9)
	assertSerieEq(t, cpy, 1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func TestEmptyCopy(t *testing.T) {
	original := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	cpy := original.EmptyCopy()
	assert.NotSame(t, original, cpy)
	assert.Equal(t, original.Type(), cpy.Type())
	assert.Equal(t, 9, original.Len())
	assert.Equal(t, 0, cpy.Len())
}
