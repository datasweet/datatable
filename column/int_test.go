package column_test

import (
	"testing"

	"github.com/datasweet/datatable/datatable/column"
	"github.com/datasweet/datatable/value"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	col := column.Int("test")
	assert.Equal(t, "test", col.Name())
	assert.Equal(t, value.Int, col.Type())
	assert.Equal(t, 32, col.NewValue().Set("32"))
}
