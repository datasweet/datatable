package datatable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	tb := New("test")
	assert.Equal(t, 0, tb.NumCols())

	tb.AddColumn("sessions", Number)
	tb.AddColumn("bounces", Number)
	tb.AddColumn("bounceRate", Number)

	cols := tb.Columns()
	assert.Len(t, cols, 3)
	assert.Equal(t, "sessions", cols[0].Name())
	assert.Equal(t, "bounces", cols[1].Name())
	assert.Equal(t, "bounceRate", cols[2].Name())
	assert.Equal(t, 0, tb.NumRows())

	tb.AddColumn("pageViews", Number, 1, 2, 3, 4, 5)
	assert.Equal(t, 4, tb.NumCols())
	assert.Equal(t, 5, tb.NumRows())

	checkTable(t, tb,
		"sessions", "bounces", "bounceRate", "pageViews",
		0.0, 0.0, 0.0, 1.0,
		0.0, 0.0, 0.0, 2.0,
		0.0, 0.0, 0.0, 3.0,
		0.0, 0.0, 0.0, 4.0,
		0.0, 0.0, 0.0, 5.0,
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

	tb.AddColumn("win", Number)
	checkTable(t, tb,
		"champ", "win",
		"Malzahar", 0.0,
		"Xerath", 0.0,
		"", 0.0,
		"Ahri", 0.0,
	)

	tb.AddColumn("loose", Number, 3, 4, nil)
	checkTable(t, tb,
		"champ", "win", "loose",
		"Malzahar", 0.0, 3.0,
		"Xerath", 0.0, 4.0,
		"", 0.0, nil,
		"Ahri", 0.0, 0.0,
	)

}

func TestExprColumn(t *testing.T) {
	tb := New("test")
	tb.AddColumn("champ", String, "Malzahar", "Xerath", "Teemo")
	tb.AddExprColumn("champion", "upper(`champ`)")
	tb.AddColumn("win", Number, 10, 20, 666)
	tb.AddColumn("loose", Number, 6, 5, 666)
	tb.AddExprColumn("winRate", "(`win` * 100 / (`win` + `loose`)) ~ \" %\"")
	tb.AddExprColumn("sum", "sum(`win`)")

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum",
		"Malzahar", "MALZAHAR", 10.0, 6.0, "62.5 %", 696.0,
		"Xerath", "XERATH", 20.0, 5.0, "80 %", 696.0,
		"Teemo", "TEEMO", 666.0, 666.0, "50 %", 696.0,
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
