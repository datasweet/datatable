package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerieFactory(t *testing.T) {
	typs := ColumnTypes()
	for _, typ := range typs {
		assert.NotPanics(t, func() { newColumnSerie(typ, ColumnOptions{}) }, typ)
	}
}
