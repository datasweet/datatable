package serie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datasweet/datatable/serie"
)

func TestSerieFloat32(t *testing.T) {
	s := serie.Float32()
	assert.NotNil(t, s)

	s.Append(1, "23", 3.14, "teemo", true, nil)

	assertSerieEq(t, s, float32(1), float32(23), float32(3.14), float32(0), float32(1), float32(0))
}

func TestSerieFloat32N(t *testing.T) {
	s := serie.Float32N()
	assert.NotNil(t, s)

	s.Append(1, "23", 3.14, "teemo", true, nil)
	assertSerieEq(t, s, float32(1), float32(23), float32(3.14), nil, float32(1), nil)
}
