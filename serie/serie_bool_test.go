package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestSerieBool(t *testing.T) {
	s := serie.Bool()
	assert.NotNil(t, s)

	s.Append(1, 0, true, "teemo", nil)

	assertSerieEq(t, s, true, false, true, false, false)

	s.SortAsc()
	assertSerieEq(t, s, false, false, false, true, true)

	s.SortDesc()
	assertSerieEq(t, s, true, true, false, false, false)

}

func TestSerieBoolN(t *testing.T) {
	s := serie.BoolN()
	assert.NotNil(t, s)

	s.Append(1, 0, true, "teemo", nil)
	assertSerieEq(t, s, true, false, true, nil, nil)

	s.SortAsc()
	assertSerieEq(t, s, nil, nil, false, true, true)

	s.SortDesc()
	assertSerieEq(t, s, true, true, false, nil, nil)

}
