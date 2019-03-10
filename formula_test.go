package datatable

import (
	"testing"
)

func TestFormula(t *testing.T) {
	tb := New("test")
	tb.AddColumn("champ", String, "Malzahar", "Xerath", "Teemo")
	tb.AddFormulaColumn("champion", String, "upper(`champ`)")
	tb.AddColumn("win", Number, 10, 20, 666)
	tb.AddColumn("loose", Number, 6, 5, 666)
	tb.AddFormulaColumn("winRate", Number, "`win` * 100 / (`win` + `loose`)")
	tb.AddFormulaColumn("sum", Number, "sum(`win`)")

	checkTable(t, tb,
		"champ", "champion", "win", "loose", "winRate", "sum",
		"Malzahar", "MALZAHAR", 10.0, 6.0, 62.5, 696.0,
		"Xerath", "XERATH", 20.0, 5.0, 80.0, 696.0,
		"Teemo", "TEEMO", 666.0, 666.0, 50.0, 696.0,
	)
}
