package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestSerieInt64(t *testing.T) {
	s := serie.Int64()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, -67, nil)
	assertSerieEq(t, s,
		int64(31),
		int64(23),
		int64(98),
		int64(0),
		int64(1),
		int64(-67),
		int64(0),
	)
}

func TestSerieInt64N(t *testing.T) {
	s := serie.Int64N()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, -67, nil)
	assertSerieEq(t, s,
		int64(31),
		int64(23),
		int64(98),
		nil,
		int64(1),
		int64(-67),
		nil,
	)
}
