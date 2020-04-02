package datatable_test

import (
	"testing"

	"github.com/datasweet/datatable/datatable"
	"github.com/datasweet/datatable/serie"
	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	tb := datatable.New("test")
	assert.Equal(t, 0, tb.NCols())

	assert.NoError(t, tb.AddColumn("sessions", serie.NewInt(120)))
	assert.NoError(t, tb.AddColumn("bounces", serie.NewInt()))
	assert.NoError(t, tb.AddColumn("bounceRate", serie.NewFloat64()))
	assert.Equal(t, 3, tb.NCols())
	assert.Equal(t, 1, tb.NRows())
	assert.NoError(t, tb.AddColumn("pageViews", serie.NewInt(1, 2, 3, 4, 5)))
	assert.Equal(t, 4, tb.NCols())
	assert.Equal(t, 5, tb.NRows())
	assert.Error(t, tb.AddColumn("nil serie", nil))
	assert.Error(t, tb.AddColumn("pageViews", serie.NewInt(1, 2, 3, 4, 5, 6, 7, 8, 9)))
	assert.Equal(t, 4, tb.NCols())
	assert.Equal(t, 5, tb.NRows())

	// cols := tb.Columns()
	// assert.Len(t, cols, 3)
	// assert.Equal(t, "sessions", cols[0].Name())
	// assert.Equal(t, "bounces", cols[1].Name())
	// assert.Equal(t, "bounceRate", cols[2].Name())
	// assert.Equal(t, 1, tb.NumRows())

	// tb.AddColumn("pageViews", Int, 1, 2, 3, 4, 5)
	// assert.Equal(t, 4, tb.NumCols())
	// assert.Equal(t, 5, tb.NumRows())

	// checkTable(t, tb,
	// 	"sessions", "bounces", "bounceRate", "pageViews",
	// 	int64(120), int64(0), 0.0, int64(1),
	// 	int64(0), int64(0), 0.0, int64(2),
	// 	int64(0), int64(0), 0.0, int64(3),
	// 	int64(0), int64(0), 0.0, int64(4),
	// 	int64(0), int64(0), 0.0, int64(5),
	// )
}
