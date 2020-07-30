package serie_test

import (
	"testing"

	"github.com/datasweet/datatable/serie"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
)

func TestNewSerie(t *testing.T) {
	assert.Panics(t, func() { serie.New(nil, nil, nil) })
	assert.Panics(t, func() { serie.New(1, nil, nil) })
	assert.Panics(t, func() {
		serie.New(1, cast.ToFloat32, func(i, j int) int {
			return serie.Eq
		})
	})
	assert.Panics(t, func() { serie.New(1, cast.ToInt, nil) })

	// OK
	s := serie.New(1, cast.ToInt, func(i, j int) int {
		return serie.Eq
	})
	assert.NotNil(t, s)
	s.Append(1, 2, 3, 4, 5)
	assertSerieEq(t, s, 1, 2, 3, 4, 5)
}
