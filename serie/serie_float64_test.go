package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestSerieFloat64(t *testing.T) {
	s := serie.Float64()
	assert.NotNil(t, s)

	s.Append(1, "23", 3.14, "teemo", true, nil)

	assertSerieEq(t, s, float64(1), float64(23), float64(3.14), float64(0), float64(1), float64(0))

	s.SortAsc()
	assertSerieEq(t, s, float64(0), float64(0), float64(1), float64(1), float64(3.14), float64(23))

	s.SortDesc()
	assertSerieEq(t, s, float64(23), float64(3.14), float64(1), float64(1), float64(0), float64(0))
}

func TestSerieFloat64N(t *testing.T) {
	s := serie.Float64N()
	assert.NotNil(t, s)

	s.Append(1, "23", 3.14, "teemo", true, nil)
	assertSerieEq(t, s, float64(1), float64(23), float64(3.14), nil, float64(1), nil)

	s.SortAsc()
	assertSerieEq(t, s, nil, nil, float64(1), float64(1), float64(3.14), float64(23))

	s.SortDesc()
	assertSerieEq(t, s, float64(23), float64(3.14), float64(1), float64(1), nil, nil)
}
