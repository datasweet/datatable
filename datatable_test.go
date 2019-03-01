package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	tb := New("test")
	assert.Len(t, tb.Cols(), 0)

	tb.AddColumn("sessions", Number)
	tb.AddColumn("bounces", Number)
	tb.AddColumn("bounceRate", Number)

	cols := tb.Cols()
	assert.Len(t, cols, 3)
	assert.Equal(t, "sessions", cols[0])
	assert.Equal(t, "bounces", cols[1])
	assert.Equal(t, "bounceRate", cols[2])
	assert.Equal(t, 0, tb.nrows)

	tb.AddColumn("pageViews", Number, 1, 2, 3, 4, 5)
	cols = tb.Cols()
	assert.Len(t, cols, 4)
	assert.Equal(t, 5, tb.nrows)

	checkTable(t, tb,
		"sessions", "bounces", "bounceRate", "pageViews",
		nil, nil, nil, float64(1),
		nil, nil, nil, float64(2),
		nil, nil, nil, float64(3),
		nil, nil, nil, float64(4),
		nil, nil, nil, float64(5),
	)
}

func checkTable(t *testing.T, tb *DataTable, cells ...interface{}) {
	ncols := len(tb.Cols())
	nrows := tb.nrows + 1 // headers !
	assert.Len(t, cells, ncols*nrows)

	raw := tb.Raw()
	assert.Len(t, raw, nrows)

	for i, v := range cells {
		r := i / ncols
		c := i % ncols
		assert.Equal(t, v, raw[r][c], "ROW #%d, COL #%d", r, c)
	}
}

// t.AddColumn("sessions", Int64)
// t.AddColumn("word", String)
// t.AddColumn("bool", Bool)

// t.AddRow(123, "abc", false)
// t.AddRow(314, "pi", true)

// t.AddColumn("bounces", Int64, 456, 789)
// t.AddColumn("bounceRate", Float64)

// t.Columns[len(t.Columns)-1].Formula = "sessions * 100 / bounces"

// t.Print()

// t.Swap(0, 2)
// t.Print()
// }
