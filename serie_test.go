package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerieFactory(t *testing.T) {
	typs := []string{
		Boolean,
		String,
		Int,
		// Int8,
		// Int16,
		Int32,
		Int64,
		// Uint,
		// Uint8,
		// Uint16,
		// Uint32,
		// Uint64,
		Float32,
		Float64,
		Time,
		Raw,
	}

	for _, typ := range typs {
		assert.NotPanics(t, func() { newSerie(typ) })
	}
}
