package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	tb := New("test")
	assert.Equal(t, 0, tb.NumCols())

	tb.AddColumn("sessions", Int, 120)
	tb.AddColumn("bounces", Int)
	tb.AddColumn("bounceRate", Float)

	cols := tb.Columns()
	assert.Len(t, cols, 3)
	assert.Equal(t, "sessions", cols[0].Name())
	assert.Equal(t, "bounces", cols[1].Name())
	assert.Equal(t, "bounceRate", cols[2].Name())
	assert.Equal(t, 1, tb.NumRows())

	tb.AddColumn("pageViews", Int, 1, 2, 3, 4, 5)
	assert.Equal(t, 4, tb.NumCols())
	assert.Equal(t, 5, tb.NumRows())

	checkTable(t, tb,
		"sessions", "bounces", "bounceRate", "pageViews",
		int64(120), int64(0), 0.0, int64(1),
		int64(0), int64(0), 0.0, int64(2),
		int64(0), int64(0), 0.0, int64(3),
		int64(0), int64(0), 0.0, int64(4),
		int64(0), int64(0), 0.0, int64(5),
	)
}

func TestNewRow(t *testing.T) {
	tb := New("test")
	tb.AddColumn("champ", String)
	assert.Equal(t, 1, tb.NumCols())

	buff := tb.NewRow().Set("champ", "Malzahar")
	assert.Equal(t, 0, tb.NumRows())
	assert.True(t, tb.AddRow(buff))
	assert.Equal(t, 1, tb.NumRows())

	assert.False(t, tb.AddRow(nil))
	assert.True(t, tb.AddRow(tb.NewRow().Set("champ", "Xerath")))
	assert.True(t, tb.AddRow(tb.NewRow().Set("satan", "Teemo"))) // wrong column => not set
	assert.True(t, tb.AddRow(tb.NewRow().Set("champ", "Ahri")))

	checkTable(t, tb,
		"champ",
		"Malzahar",
		"Xerath",
		"",
		"Ahri",
	)

	tb.AddColumn("win", Int)
	checkTable(t, tb,
		"champ", "win",
		"Malzahar", int64(0),
		"Xerath", int64(0),
		"", int64(0),
		"Ahri", int64(0),
	)

	tb.AddColumn("loose", Int, 3, 4, nil)
	checkTable(t, tb,
		"champ", "win", "loose",
		"Malzahar", int64(0), int64(3),
		"Xerath", int64(0), int64(4),
		"", int64(0), nil,
		"Ahri", int64(0), int64(0),
	)

}

func TestExprColumn(t *testing.T) {
	tb := New("test")
	tb.AddColumn("champ", String, "Malzahar", "Xerath", "Teemo")
	tb.AddExprColumn("champion", "upper(`champ`)")
	tb.AddColumn("win", Int, 10, 20, 666)
	tb.AddColumn("loose", Int, 6, 5, 666)
	tb.AddExprColumn("winRate", "(`win` * 100 / (`win` + `loose`)) ~ \" %\"")
	tb.AddExprColumn("sum", "sum(`win`)")

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum",
		"Malzahar", "MALZAHAR", int64(10), int64(6), "62.5 %", 696.0,
		"Xerath", "XERATH", int64(20), int64(5), "80 %", 696.0,
		"Teemo", "TEEMO", int64(666), int64(666), "50 %", 696.0,
	)
}

func TestAppendRow(t *testing.T) {
	tb := New("test")
	tb.AddColumn("champ", String)
	tb.AddColumn("win", Int)
	tb.AddColumn("loose", Int)
	tb.AddExprColumn("winRate", "(`win` * 100 / (`win` + `loose`))")

	assert.True(t, tb.AppendRow("Xerath", 25, 15, "expr"))
	assert.True(t, tb.AppendRow("Malzahar", 16, 16, nil))
	assert.True(t, tb.AppendRow("Vel'Koz", 7, 5, 3))

	checkTable(t, tb,
		"champ", "win", "loose", "winRate",
		"Xerath", int64(25), int64(15), 62.5,
		"Malzahar", int64(16), int64(16), 50.0,
		"Vel'Koz", int64(7), int64(5), 58.333333333333336,
	)
}

func TestSwap(t *testing.T) {
	tb := New("test")
	tb.AddColumn("id", String, "0001", "0002", "0003", "0004", "0005", "0006")
	tb.AddColumn("type", String, "donut", "donut", "donut", "bar", "twist", "filled")
	tb.AddColumn("name", String, "Cake", "Raised", "Old Fashioned", "Bar", "Twist", "Filled")
	checkTable(t, tb,
		"id", "type", "name",
		"0001", "donut", "Cake",
		"0002", "donut", "Raised",
		"0003", "donut", "Old Fashioned",
		"0004", "bar", "Bar",
		"0005", "twist", "Twist",
		"0006", "filled", "Filled",
	)

	assert.False(t, tb.Swap("toto", "name"))
	checkTable(t, tb,
		"id", "type", "name",
		"0001", "donut", "Cake",
		"0002", "donut", "Raised",
		"0003", "donut", "Old Fashioned",
		"0004", "bar", "Bar",
		"0005", "twist", "Twist",
		"0006", "filled", "Filled",
	)

	assert.False(t, tb.Swap("name", "toto"))
	checkTable(t, tb,
		"id", "type", "name",
		"0001", "donut", "Cake",
		"0002", "donut", "Raised",
		"0003", "donut", "Old Fashioned",
		"0004", "bar", "Bar",
		"0005", "twist", "Twist",
		"0006", "filled", "Filled",
	)

	assert.True(t, tb.Swap("id", "name"))
	checkTable(t, tb,
		"name", "type", "id",
		"Cake", "donut", "0001",
		"Raised", "donut", "0002",
		"Old Fashioned", "donut", "0003",
		"Bar", "bar", "0004",
		"Twist", "twist", "0005",
		"Filled", "filled", "0006",
	)
}
