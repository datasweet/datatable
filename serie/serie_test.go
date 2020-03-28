package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func TestNewInt64Serie(t *testing.T) {
	s := serie.NewInt64()
	assert.NotNil(t, s)
	assert.Equal(t, serie.Int64, s.Type())
	assert.Equal(t, 0, s.Len())
	assertSerieEq(t, s, "nil")

	s.Append(1, 2, 3, 4, 5, nil, "10", "teemo", true)
	assert.Equal(t, 9, s.Len())
	assertSerieEq(t, s, "1 2 3 4 5 #NULL! 10 #NULL! 1")

	s.Prepend(-1, -2, -3, -4, "-10", "teemo", false)
	assert.Equal(t, 16, s.Len())
	assertSerieEq(t, s, "-1 -2 -3 -4 -10 #NULL! 0 1 2 3 4 5 #NULL! 10 #NULL! 1")

	assert.NoError(t, s.Insert(7, 100, 101, 102))
	assert.Equal(t, 19, s.Len())
	assertSerieEq(t, s, "-1 -2 -3 -4 -10 #NULL! 100 101 102 0 1 2 3 4 5 #NULL! 10 #NULL! 1")
}
