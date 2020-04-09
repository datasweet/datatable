package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/datasweet/datatable/value"
	"github.com/stretchr/testify/assert"
)

func TestNewIntSerie(t *testing.T) {
	s := serie.Int()
	assert.NotNil(t, s)
	assert.Equal(t, value.IntType, s.Type())
	assert.Equal(t, 0, s.Len())
	assertSerieEq(t, s)

	s.Append(1, 2, 3, 4, 5, nil, "10", "teemo", true)
	assert.Equal(t, 9, s.Len())
	assertSerieEq(t, s, 1, 2, 3, 4, 5, nil, 10, nil, 1)

	s.Prepend(-1, -2, -3, -4, "-10", "teemo", false)
	assert.Equal(t, 16, s.Len())
	assertSerieEq(t, s, -1, -2, -3, -4, -10, nil, 0, 1, 2, 3, 4, 5, nil, 10, nil, 1)

	s.Insert(6, 100, 101, 102)
	assert.Equal(t, 19, s.Len())
	assertSerieEq(t, s, -1, -2, -3, -4, -10, nil, 100, 101, 102, 0, 1, 2, 3, 4, 5, nil, 10, nil, 1)
}
