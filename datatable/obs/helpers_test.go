package datatable_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// // checkTable to check if a table contains cells
// func checkTable(t *testing.T, tb DataTable, cells ...interface{}) {
// 	var ncols int
// 	for _, c := range tb.Columns() {
// 		if !c.Hidden() {
// 			ncols++
// 		}
// 	}
// 	nrows := tb.NumRows() + 1 // headers !
// 	assert.Len(t, cells, ncols*nrows)

// 	raw := ToTable(tb, true)
// 	assert.Len(t, raw, nrows)

// 	for i, v := range cells {
// 		r := i / ncols
// 		c := i % ncols
// 		assert.Equal(t, v, raw[r][c], "ROW #%d, COL #%d", r, c)
// 	}
// }
