package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestSerieInt32(t *testing.T) {
	s := serie.Int32()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, -67, nil)
	assertSerieEq(t, s,
		int32(31),
		int32(23),
		int32(98),
		int32(0),
		int32(1),
		int32(-67),
		int32(0),
	)

	s.SortAsc()
	assertSerieEq(t, s,
		int32(-67),
		int32(0),
		int32(0),
		int32(1),
		int32(23),
		int32(31),
		int32(98),
	)

	s.SortDesc()
	assertSerieEq(t, s,
		int32(98),
		int32(31),
		int32(23),
		int32(1),
		int32(0),
		int32(0),
		int32(-67),
	)
}

func TestSerieInt32N(t *testing.T) {
	s := serie.Int32N()
	assert.NotNil(t, s)

	s.Append(31, "23", 98.5, "teemo", true, -67, nil)
	assertSerieEq(t, s,
		int32(31),
		int32(23),
		int32(98),
		nil,
		int32(1),
		int32(-67),
		nil,
	)

	s.SortAsc()
	assertSerieEq(t, s,
		nil,
		nil,
		int32(-67),
		int32(1),
		int32(23),
		int32(31),
		int32(98),
	)

	s.SortDesc()
	assertSerieEq(t, s,
		int32(98),
		int32(31),
		int32(23),
		int32(1),
		int32(-67),
		nil,
		nil,
	)
}
