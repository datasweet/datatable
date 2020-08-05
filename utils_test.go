package datatable_test

import (
	"testing"

	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

// checkTable to check if a table contains cells
func checkTable(t *testing.T, tb *datatable.DataTable, cells ...interface{}) {
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

func New(t *testing.T) *datatable.DataTable {
	tb := datatable.New("test")
	tb.AddColumn("champ", datatable.String, datatable.ColumnValues("Malzahar", "Xerath", "Teemo"))
	tb.AddColumn("champion", datatable.String, datatable.ColumnExpr("upper(`champ`)"))
	tb.AddColumn("win", datatable.Int, datatable.ColumnValues(10, 20, 666))
	tb.AddColumn("loose", datatable.Int, datatable.ColumnValues(6, 5, 666))
	tb.AddColumn("winRate", datatable.String, datatable.ColumnExpr("(`win` * 100 / (`win` + `loose`)) ~ \" %\""))
	tb.AddColumn("sum", datatable.Float64, datatable.ColumnExpr("sum(`win`)"))
	tb.AddColumn("ok", datatable.Bool, datatable.ColumnExpr("true"))
	tb.AddColumn("hidden", datatable.Bool, datatable.ColumnExpr("false"))
	tb.HideColumn("hidden")

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum", "ok",
		"Malzahar", "MALZAHAR", 10, 6, "62.5 %", 696.0, true,
		"Xerath", "XERATH", 20, 5, "80 %", 696.0, true,
		"Teemo", "TEEMO", 666, 666, "50 %", 696.0, true,
	)

	return tb
}
