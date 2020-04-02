package datatable_test

import (
	"testing"

	"github.com/datasweet/datatable/datatable"
	"github.com/stretchr/testify/assert"
)

// checkTable to check if a table contains cells
func checkTable(t *testing.T, tb datatable.DataTable, cells ...interface{}) {
	ncols := tb.NumCols()
	nrows := tb.NumRows()
	assert.Len(t, cells, ncols*(nrows+1)) // + headers

	cols := tb.Columns()
	rows := tb.Rows()

	for i, v := range cells {
		r := i/ncols - 1
		c := i % ncols

		if r == -1 {
			assert.Equal(t, v, cols[c], "HEADER COL #%d", r, c)
			continue
		}

		assert.Equal(t, v, rows[r][cols[c]], "ROW #%d, COL #%d", r, c)
	}
}
