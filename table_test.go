package datatable_test

import (
	"testing"

	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	tb := datatable.New("test")
	assert.Equal(t, 0, tb.NumCols())

	assert.NoError(t, tb.AddColumn("sessions", datatable.Int, 120))
	assert.NoError(t, tb.AddColumn("bounces", datatable.Int))
	assert.NoError(t, tb.AddColumn("bounceRate", datatable.Float64))
	assert.Error(t, tb.AddColumn("bounces", datatable.Int, 11))
	assert.Error(t, tb.AddColumn("    ", datatable.Int, 11))
	assert.Error(t, tb.AddColumn("nil", datatable.ColumnType("unknown")))

	assert.Equal(t, []string{"sessions", "bounces", "bounceRate"}, tb.Columns())
	assert.Equal(t, 1, tb.NumRows())

	assert.NoError(t, tb.AddColumn("pageViews", datatable.Int, 1, 2, 3, 4, 5))
	assert.Equal(t, 4, tb.NumCols())
	assert.Equal(t, 5, tb.NumRows())

	checkTable(t, tb,
		"sessions", "bounces", "bounceRate", "pageViews",
		120, 0, 0.0, 1,
		0, 0, 0.0, 2,
		0, 0, 0.0, 3,
		0, 0, 0.0, 4,
		0, 0, 0.0, 5,
	)
}

func TestNewRow(t *testing.T) {
	tb := datatable.New("test")
	assert.NoError(t, tb.AddColumn("champ", datatable.String))
	assert.Equal(t, 1, tb.NumCols())
	assert.Equal(t, 0, tb.NumRows())

	r := make(datatable.Row)
	r["champ"] = "Malzahar"
	tb.Append(r)
	assert.Equal(t, 1, tb.NumRows())

	tb.Append(nil)
	assert.Equal(t, 1, tb.NumRows())

	tb.Append()
	assert.Equal(t, 1, tb.NumRows())

	tb.Append(
		tb.NewRow().Set("champ", "Xerath"),
		tb.NewRow().Set("satan", "Teemo"), // wrong column => not set
		tb.NewRow().Set("champ", "Ahri"),
	)

	checkTable(t, tb,
		"champ",
		"Malzahar",
		"Xerath",
		"",
		"Ahri",
	)

	tb.AddColumn("win", datatable.Int)
	checkTable(t, tb,
		"champ", "win",
		"Malzahar", 0,
		"Xerath", 0,
		"", 0,
		"Ahri", 0,
	)

	tb.AddColumn("loose", datatable.Int, 3, 4, nil)
	checkTable(t, tb,
		"champ", "win", "loose",
		"Malzahar", 0, 3,
		"Xerath", 0, 4,
		"", 0, 0,
		"Ahri", 0, 0,
	)
}

func TestExprColumn(t *testing.T) {
	tb := datatable.New("test")
	tb.AddColumn("champ", datatable.String, "Malzahar", "Xerath", "Teemo")
	tb.AddExprColumn("champion", datatable.String, "upper(`champ`)")
	tb.AddColumn("win", datatable.Int, 10, 20, 666)
	tb.AddColumn("loose", datatable.Int, 6, 5, 666)
	tb.AddExprColumn("winRate", datatable.String, "(`win` * 100 / (`win` + `loose`)) ~ \" %\"")
	tb.AddExprColumn("sum", datatable.Int, "sum(`win`)")
	tb.AddExprColumn("ok", datatable.Bool, "true")

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum", "ok",
		"Malzahar", "MALZAHAR", 10, 6, "62.5 %", 696, true,
		"Xerath", "XERATH", 20, 5, "80 %", 696, true,
		"Teemo", "TEEMO", 666, 666, "50 %", 696, true,
	)
}

func TestAppendRow(t *testing.T) {
	tb := datatable.New("test")
	assert.NoError(t, tb.AddColumn("champ", datatable.String))
	assert.NoError(t, tb.AddColumn("win", datatable.Int))
	assert.NoError(t, tb.AddColumn("loose", datatable.Int))
	assert.NoError(t, tb.AddExprColumn("winRate", datatable.Float64, "(`win` * 100 / (`win` + `loose`))"))
	assert.Error(t, tb.AddExprColumn("winRate", datatable.String, "test"))

	assert.NoError(t, tb.AppendRow("Xerath", 25, 15, "expr"))
	assert.NoError(t, tb.AppendRow("Malzahar", 16, 16, nil))
	assert.NoError(t, tb.AppendRow("Vel'Koz", 7, 5, 3))

	checkTable(t, tb,
		"champ", "win", "loose", "winRate",
		"Xerath", 25, 15, 62.5,
		"Malzahar", 16, 16, 50.0,
		"Vel'Koz", 7, 5, 58.333333333333336,
	)
}
