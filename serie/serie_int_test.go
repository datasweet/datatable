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
}

func TestSerieIntN(t *testing.T) {
	s := serie.IntN()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, -67, nil)
	assertSerieEq(t, s, 31, 23, 98, nil, 1, -67, nil)
}
