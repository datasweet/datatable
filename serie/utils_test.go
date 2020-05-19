package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func assertSerieEq(t *testing.T, s serie.Serie, v ...interface{}) {
	assert.NotNil(t, s)
	assert.Equal(t, len(v), s.Len())
	for i, val := range s.All() {
		assert.Equalf(t, v[i], val, "At index %d", i)
	}
}
