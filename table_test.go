package datatable_test

import (
	"fmt"
	"testing"

	"github.com/datasweet/datatable"
	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	tb := datatable.New("test")
	assert.Equal(t, 0, tb.NumCols())

	assert.NoError(t, tb.AddColumn("sessions", datatable.Int, datatable.Values(120)))
	assert.NoError(t, tb.AddColumn("bounces", datatable.Int))
	assert.NoError(t, tb.AddColumn("bounceRate", datatable.Float64))
	assert.Error(t, tb.AddColumn("bounces", datatable.Int, datatable.Values(11)))
	assert.Error(t, tb.AddColumn("    ", datatable.Int, datatable.Values(11)))
	assert.Error(t, tb.AddColumn("nil", datatable.ColumnType("unknown")))

	assert.Equal(t, []string{"sessions", "bounces", "bounceRate"}, tb.Columns())
	assert.Equal(t, 1, tb.NumRows())

	assert.NoError(t, tb.AddColumn("pageViews", datatable.Int, datatable.Values(1, 2, 3, 4, 5)))
	assert.Equal(t, 4, tb.NumCols())
	assert.Equal(t, 5, tb.NumRows())

	fmt.Println(tb)

	checkTable(t, tb,
		"sessions", "bounces", "bounceRate", "pageViews",
		120, nil, nil, 1,
		nil, nil, nil, 2,
		nil, nil, nil, 3,
		nil, nil, nil, 4,
		nil, nil, nil, 5,
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
		nil,
		"Ahri",
	)

	tb.AddColumn("win", datatable.Int)
	checkTable(t, tb,
		"champ", "win",
		"Malzahar", nil,
		"Xerath", nil,
		nil, nil,
		"Ahri", nil,
	)

	tb.AddColumn("loose", datatable.Int, datatable.Values(3, 4, nil))
	checkTable(t, tb,
		"champ", "win", "loose",
		"Malzahar", nil, 3,
		"Xerath", nil, 4,
		nil, nil, nil,
		"Ahri", nil, nil,
	)
}

func TestExprColumn(t *testing.T) {
	tb := datatable.New("test")
	tb.AddColumn("champ", datatable.String, datatable.Values("Malzahar", "Xerath", "Teemo"))
	tb.AddColumn("champion", datatable.String, datatable.Expr("upper(`champ`)"))
	tb.AddColumn("win", datatable.Int, datatable.Values(10, 20, 666))
	tb.AddColumn("loose", datatable.Int, datatable.Values(6, 5, 666))
	tb.AddColumn("winRate", datatable.String, datatable.Expr("(`win` * 100 / (`win` + `loose`)) ~ \" %\""))
	tb.AddColumn("sum", datatable.Int, datatable.Expr("sum(`win`)"))
	tb.AddColumn("ok", datatable.Bool, datatable.Expr("true"))

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
	assert.NoError(t, tb.AddColumn("winRate", datatable.Float64, datatable.Expr("(`win` * 100 / (`win` + `loose`))")))
	assert.Error(t, tb.AddColumn("winRate", datatable.String, datatable.Expr("test")))

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
