package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestSerieInt(t *testing.T) {
	s := serie.Int()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, -67, nil)
	assertSerieEq(t, s, 31, 23, 98, 0, 1, -67, 0)

	s.SortAsc()
	assertSerieEq(t, s, -67, 0, 0, 1, 23, 31, 98)

	s.SortDesc()
	assertSerieEq(t, s, 98, 31, 23, 1, 0, 0, -67)
}

func TestSerieIntN(t *testing.T) {
	s := serie.IntN()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, -67, nil)
	assertSerieEq(t, s, 31, 23, 98, nil, 1, -67, nil)

	s.SortAsc()
	assertSerieEq(t, s, nil, nil, -67, 1, 23, 31, 98)

	s.SortDesc()
	assertSerieEq(t, s, 98, 31, 23, 1, -67, nil, nil)
}
