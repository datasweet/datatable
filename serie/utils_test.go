package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func assertSerieEq(t *testing.T, s serie.Serie, v ...interface{}) {
	assert.NotNil(t, s)
	assert.Equal(t, len(v), s.Len())
	index := 0
	for it := s.Iterator(); it.Next(); {
		assert.Equalf(t, v[index], it.Current(), "At index %d", index)
		index++
	}
}
