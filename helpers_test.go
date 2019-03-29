package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// rowsEq to test if a column contains row values.
func rowsEq(t *testing.T, col *column, values ...interface{}) {
	assert.Equal(t, len(values), col.Len(), "Len() failed: %v", col.Rows())

	for i, v := range values {
		assert.Equal(t, v, col.GetAt(i), "Values() failed: %v", col.Rows())
	}
}

// checkTable to check if a table contains cells
func checkTable(t *testing.T, tb DataTable, cells ...interface{}) {
	ncols := tb.NumCols()
	nrows := tb.NumRows() + 1 // headers !
	assert.Len(t, cells, ncols*nrows)

	raw := ToTable(tb, true)
	assert.Len(t, raw, nrows)

	for i, v := range cells {
		r := i / ncols
		c := i % ncols
		assert.Equal(t, v, raw[r][c], "ROW #%d, COL #%d", r, c)
	}
}
