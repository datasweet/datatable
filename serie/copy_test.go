package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"

	"github.com/stretchr/testify/assert"
)

func TestShallowCopy(t *testing.T) {
	original := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	shallow := original.Copy(serie.ShallowCopy)
	assert.NotSame(t, original, shallow)
	assert.Equal(t, original.Type(), shallow.Type())
	assert.Equal(t, original.Len(), shallow.Len())
	assertSerieEq(t, shallow, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Shallow should share same value
	original.Value(4).Set(50)
	assertSerieEq(t, original, 1, 2, 3, 4, 50, 6, 7, 8, 9)
	assertSerieEq(t, shallow, 1, 2, 3, 4, 50, 6, 7, 8, 9)

	// but not the same values
	original.Append(10)
	assertSerieEq(t, original, 1, 2, 3, 4, 50, 6, 7, 8, 9, 10)
	assertSerieEq(t, shallow, 1, 2, 3, 4, 50, 6, 7, 8, 9)
}

func TestDeepCopy(t *testing.T) {
	original := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	shallow := original.Copy(serie.DeepCopy)
	assert.NotSame(t, original, shallow)
	assert.Equal(t, original.Type(), shallow.Type())
	assert.Equal(t, original.Len(), shallow.Len())
	assertSerieEq(t, shallow, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Deep should not share same value
	original.Value(4).Set(50)
	assertSerieEq(t, original, 1, 2, 3, 4, 50, 6, 7, 8, 9)
	assertSerieEq(t, shallow, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	original.Append(10)
	assertSerieEq(t, original, 1, 2, 3, 4, 50, 6, 7, 8, 9, 10)
	assertSerieEq(t, shallow, 1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func TestEmptyCopy(t *testing.T) {
	original := serie.Int(1, 2, 3, 4, 5, 6, 7, 8, 9)
	shallow := original.Copy(serie.EmptyCopy)
	assert.NotSame(t, original, shallow)
	assert.Equal(t, original.Type(), shallow.Type())
	assert.Equal(t, 0, shallow.Len())
	assertSerieEq(t, shallow)
}
