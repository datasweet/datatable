package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestSerieRaw(t *testing.T) {
	s := serie.Raw()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, nil, -67)
	assertSerieEq(t, s, 31, "23", 98.5, "teemo", true, nil, -67)

	s.SortAsc()
	assertSerieEq(t, s, nil, -67, "23", 31, 98.5, "teemo", true)

	s.SortDesc()
	assertSerieEq(t, s, true, "teemo", 98.5, 31, "23", -67, nil)
}
